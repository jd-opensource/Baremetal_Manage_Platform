from pypureomapi import (Omapi)


class OmapiClient(object):
    def __init__(self, omapi_key, host='127.0.0.1', port=7911):
        self._omapi = Omapi(
            host,
            int(port),
            b"omapi_key",
            omapi_key.encode("ascii"),
        )

    def add_host(self, ip, mac):
        name = self._name_from_mac(mac)
        self._omapi.add_host_supersede(ip, mac, name)

    def delete_host(self, mac):
        self._omapi.del_host(mac)

    def exists_host(self, mac):
        try:
            self._omapi.lookup_host_host(mac)
            return True
        except Exception:
            return False

    def _name_from_mac(self, mac):
        return mac.replace(":", "-").encode("ascii")
