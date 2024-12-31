import logging
import os
import subprocess
import time

from bmpda.utils import encode

LOG = logging.getLogger(__name__)


class UnknownArgumentError(Exception):

    def __init__(self, message=None):
        super(UnknownArgumentError, self).__init__(message)


class ProcessExecutionError(Exception):

    def __init__(self,
                 stdout=None,
                 stderr=None,
                 exit_code=None,
                 cmd=None,
                 description=None):
        super(ProcessExecutionError, self).__init__(stdout, stderr, exit_code,
                                                    cmd, description)
        self.stdout = stdout
        self.stderr = stderr
        self.exit_code = exit_code
        self.cmd = cmd
        self.description = description

    def __str__(self):
        description = self.description
        if description is None:
            description = "Unexpected error while running command."

        exit_code = self.exit_code
        if exit_code is None:
            exit_code = '-'

        message_param = {
            'description': description,
            'cmd': self.cmd,
            'exit_code': exit_code,
            'stdout': self.stdout
        }
        message = ('%(description)s\n'
                   'Command: %(cmd)s\n'
                   'Exit code: %(exit_code)s\n'
                   'Stdout: %(stdout)r') % message_param
        return message


class NoRootPermissionError(ProcessExecutionError):

    def __init__(self, **kwargs):
        super(NoRootPermissionError, self).__init__(**kwargs)


def try_execute(*cmd, **kwargs):
    """The same as execute but returns None on error.

    Instead of raising an exception on failure, this method simply
    returns None in case of failure.

    :param *cmd: positional arguments to pass to execute()
    :param **kwargs: keyword arguments to pass to processutils.execute()
    :raises: UnknownArgumentException on receiving unknown arguments
    :returns: tuple of (stdout, stderr) or None in some error cases
    """
    try:
        return execute(*cmd, **kwargs)
    except (ProcessExecutionError, OSError) as e:
        LOG.info('Command failed: %s', e)


def execute(*cmd, **kwargs):
    """Helper method to shell out and execute a command through subprocess.

    Allows optional.

    :param cmd:             Passed to subprocess.Popen.
    :type cmd:              string

    :param cwd:             Set the current working directory
    :type cwd:              string

    :param target:          Chroot dir
    :type target:           string

    :type std_input:        string or bytes

    :param env_variables:   Environment variables and their values that
                            will be set for the process.
    :type env_variables:    dict

    :param check_exit_code: Single bool, int, or list of allowed exit
                            codes.  Defaults to [0].  Raise
                            :class:`ProcessExecutionException` unless
                            program exits with one of these code.
    :type check_exit_code:  boolean, int, or [int]


    :param delay_on_retry:  True | False. Defaults to True. If set to True,
                            wait a short amount of time before retrying.
    :type delay_on_retry:   boolean

    :param attempts_times:  How many times to retry cmd.
    :type attempts_times:   int

    :param run_as_root:     True | False. Defaults to False. If set to True,
                            the command is prefixed by the command specified
                            in the root_helper kwarg.
    :type run_as_root:      boolean

    :param shell:           whether or not there should be a shell used to
                            execute this command. Defaults to True.
    :type shell:            boolean

    :param binary:          On Python 3, return stdout and stderr as bytes if
                            binary is True, as Unicode otherwise.
    :type binary:           boolean

    :param use_standard_locale: keyword-only argument. True | False.
                            Defaults to True. If set to True,
                            execute command with standard locale
                            added to environment variables.
    :type                   boolean

    :returns:               (status, stdout, stderr) from process execution
    """

    target = kwargs.pop('target', '/')
    chroot_args = [] if target == '/' else ['chroot', target]
    cmd = [str(c) for c in cmd]
    cmd = chroot_args + cmd
    cmd_text = ' '.join(cmd)
    use_standard_locale = kwargs.pop('use_standard_locale', True)
    if use_standard_locale:
        environ = os.environ.copy()
        environ['PATH'] = '%s:/bin' % environ['PATH']
        env = kwargs.pop('env_variables', environ)
        env['LC_ALL'] = 'C'
        kwargs['env_variables'] = env

    cwd = kwargs.pop('cwd', None)

    std_input = kwargs.pop('std_input', None)
    if std_input is not None:
        std_input = encode.to_utf8(std_input)

    env_variables = kwargs.pop('env_variables', None)

    ignore_exit_code = False
    check_exit_code = kwargs.pop('check_exit_code', [0])
    if isinstance(check_exit_code, bool):
        ignore_exit_code = not check_exit_code
        check_exit_code = [0]
    elif isinstance(check_exit_code, int):
        check_exit_code = [check_exit_code]

    delay_on_retry = kwargs.pop('delay_on_retry', True)

    attempts_times = kwargs.pop('attempts_times', 1)

    run_as_root = kwargs.pop('run_as_root', False)
    if run_as_root and hasattr(os, 'geteuid') and os.geteuid() != 0:
        raise NoRootPermissionError(
            cmd=cmd_text,
            description='Command requested root, but did not '
            'specify a root helper.')

    shell = kwargs.pop('shell', False)

    binary = kwargs.pop('binary', False)

    if kwargs:
        raise UnknownArgumentError('Got unknown keyword args: %r' % kwargs)

    temp_attempts_times = attempts_times
    while temp_attempts_times > 0:
        temp_attempts_times -= 1

        LOG.info("ready to run %s", cmd_text)
        try:
            _PIPE = subprocess.PIPE
            obj = subprocess.Popen(cmd,
                                   stdin=_PIPE,
                                   stdout=_PIPE,
                                   stderr=subprocess.STDOUT,
                                   close_fds=True,
                                   preexec_fn=None,
                                   shell=shell,
                                   cwd=cwd,
                                   env=env_variables)
            result = obj.communicate(std_input)
            obj.stdin.close()
            _returncode = obj.returncode
            if not ignore_exit_code and _returncode not in check_exit_code:
                (stdout, stderr) = result
                stdout = os.fsdecode(stdout) if stdout is not None else None
                stderr = os.fsdecode(stderr) if stderr is not None else None
                raise ProcessExecutionError(exit_code=_returncode,
                                            stdout=stdout,
                                            stderr=stderr,
                                            cmd=cmd)
            if not binary and result is not None:
                (stdout, stderr) = result
                # Decode from the locale using using the surrogateescape error
                # handler (decoding cannot fail)
                stdout = os.fsdecode(stdout) if stdout is not None else None
                stderr = os.fsdecode(stderr) if stderr is not None else None
                return (stdout, stderr)
            else:
                return result
        except (ProcessExecutionError, OSError) as e:
            if isinstance(e, ProcessExecutionError):
                LOG.error(
                    "%(desc)r\ncommand: %(cmd)r\nexit code: %(code)r\n' \
                    stdout: %(stdout)r\n" % {
                        "desc": e.description,
                        "cmd": e.cmd,
                        "code": e.exit_code,
                        "stdout": e.stdout
                    })
            else:
                LOG.error(
                    'Got an OSError, command: %(cmd)r, errno: %(errno)r' % {
                        "cmd": cmd_text,
                        "errno": e.strerror
                    })
            if delay_on_retry and temp_attempts_times > 0:
                attempts_period = 0.5 * 2**(attempts_times -
                                            temp_attempts_times)
                LOG.info("retry to run `%s` after %.2f seconds" %
                         (cmd, attempts_period))
                time.sleep(attempts_period)
                continue
            raise

        finally:
            # this appears to be necessary to let the subprocess
            # call clean something up in between calls, without
            # it two execute calls in a row hangs the second one
            # Above is probably specific to the eventlet subprocess
            # module eventlet subprocess module, but since we still
            # have to support that we're leaving the sleep.  It won't
            # hurt anything in the stdlib case anyway.
            time.sleep(0)
