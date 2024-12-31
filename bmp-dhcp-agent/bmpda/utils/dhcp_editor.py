import re


class DHCPEditor(object):
    config_file = ''

    def __init__(self, config_file_path):
        self.config_file = config_file_path

    def add_host(self, hostname, ip, mac):
        with open(self.config_file, 'a') as dhcp_list:
            host_body = ' {\n\thardware ethernet ' + mac + \
                ';\n\tfixed-address ' + ip + '; \n'
            dhcp_list.write('host ' + hostname + host_body)

    def delete_host(self, mac):
        """Can delete entries either by mac.

        IT CANNOT DELETE ANYTHING BY HOSTNAME.
        """
        mac_format = re.compile(r'(.*)hardware(\s+)ethernet(\s+)%s;' % mac)
        intext = []
        with open(self.config_file, 'r') as dhcp_list:
            intext = dhcp_list.readlines()
            wait_delete_line_num = []
            for i in range(len(intext)):
                if mac_format.match(intext[i]) is not None:
                    for j in range(i, len(intext)):
                        if intext[j].find('}') != -1:
                            wait_delete_line_num.append(j)
                            break

                        wait_delete_line_num.append(j)

                    for k in range(i, -1, -1):
                        if intext[k].find('{') != -1:
                            wait_delete_line_num.append(k)
                            break

                        wait_delete_line_num.append(k)

        with open(self.config_file, 'w') as dhcp_file:
            for i in range(len(intext)):
                if i not in wait_delete_line_num:
                    dhcp_file.write(intext[i])

    def exists_host(self, mac):  # search by mac
        with open(self.config_file, 'r') as dhcp_list:
            mac_format = re.compile(r'(.*)hardware(\s+)ethernet(\s+)%s;' % mac)
            intext = dhcp_list.readlines()
            for i in range(len(intext)):
                if mac_format.match(intext[i]) is not None:
                    return True
        return False

    def exists_hostname(self, hostname):  # search by hostname
        with open(self.config_file, 'r') as dhcp_list:
            intext = dhcp_list.readlines()
            for i in range(len(intext)):
                if hostname in intext[i]:
                    return True
        return False

    def exists_subnet(self, subnet):
        with open(self.config_file, 'r') as dhcp_list:
            pattern = re.compile(r'^[\s]*subnet(\s+)%s(\s+)netmask' % subnet)
            intext = dhcp_list.readlines()
            for i in range(len(intext)):
                if pattern.match(intext[i]) is not None:
                    return True
        return False

    def add_subnet(self, subnet, subnet_mask, routes):
        with open(self.config_file, 'a') as dhcp_list:
            subnet_body = ' {\n\toption subnet-mask ' + subnet_mask + \
                ';\n\toption routers ' + routes + '; \n}\n'
            dhcp_list.write(
                'subnet %(subnet)s netmask %(subnet_mask)s %(subnet_body)s' % {
                    'subnet': subnet,
                    'subnet_mask': subnet_mask,
                    'subnet_body': subnet_body
                })

    def delete_subnet(self, subnet):
        """Can delete entries either by subnet.

        subnet x.x.x.x netmask y.y.y.y {
        }
        """
        pattern = re.compile(r'^[\s]*subnet(\s+)%s(\s+)netmask' % subnet)
        self.delete_obj_by_pattern(pattern)

    def delete_obj_by_pattern(self, pattern):
        """Delete obj by pattern

        obj:
        xxx {

        }
        """
        intext = []
        with open(self.config_file, 'r') as dhcp_list:
            intext = dhcp_list.readlines()
            wait_delete_line_num = []
            for i in range(len(intext)):
                if pattern.match(intext[i]) is not None:
                    for j in range(i, len(intext)):
                        if intext[j].find('}') != -1:
                            wait_delete_line_num.append(j)
                            break

                        wait_delete_line_num.append(j)

                    for k in range(i, -1, -1):
                        if intext[k].find('{') != -1:
                            wait_delete_line_num.append(k)
                            break

                        wait_delete_line_num.append(k)

        with open(self.config_file, 'w') as dhcp_file:
            for i in range(len(intext)):
                if i not in wait_delete_line_num:
                    dhcp_file.write(intext[i])
