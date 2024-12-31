{% for mount in mounts %}
LABEL={{mount.label}}    {{mount.mountpoint}}     {{mount.fs_type}}    {{mount.options}}   0 0
{% endfor %}