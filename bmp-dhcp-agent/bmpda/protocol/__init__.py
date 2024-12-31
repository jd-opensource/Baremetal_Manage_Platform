from abc import ABC, abstractmethod

from bmpda.command.common import Command


class BaseProtocol(ABC):

    @abstractmethod
    def send_command(self, command: Command):
        raise NotImplementedError()

    @abstractmethod
    def receive_command(self, command: str):
        raise NotImplementedError()


class Protocol(BaseProtocol):

    def start():
        pass

    def stop(self):
        pass
