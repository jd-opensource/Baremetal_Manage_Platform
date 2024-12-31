package util

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

// SSHClient SSH client
type SSHClient struct {
	Host     string
	Port     int
	Username string
	Password string
	client   *ssh.Client
	session  *ssh.Session
}

// Exec run command at remote ssh
func (s *SSHClient) Exec(cmd string, params ...interface{}) (res string, err error) {

	var (
		buf string
		ret []byte
	)
	for _, arg := range params {
		buf = fmt.Sprintf("%s %s", buf, arg)
	}

	cmdStr := cmd + " " + buf

	fmt.Println(cmdStr)

	err = s.connect()
	fmt.Println(s.session)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	ret, err = s.session.CombinedOutput(cmdStr)

	defer func() {
		s.session.Close()
		s.client.Close()
	}()

	return string(ret), err

}

func (s *SSHClient) connect() error {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		config       ssh.Config
		err          error
	)

	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(s.Password))

	config = ssh.Config{
		Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
	}

	clientConfig = &ssh.ClientConfig{
		User:    s.Username,
		Auth:    auth,
		Timeout: 300 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connect to ssh
	addr = fmt.Sprintf("%s:%d", s.Host, s.Port)

	if s.client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		fmt.Println(err.Error())
		return err
	}

	// create session
	if s.session, err = s.client.NewSession(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := s.session.RequestPty("xterm", 80, 40, modes); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
