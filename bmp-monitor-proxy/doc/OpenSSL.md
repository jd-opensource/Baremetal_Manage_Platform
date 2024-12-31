
# OpenSSL证书生成

[TOC]

## 1. **CA根证书**

* ca.crt
```bash
openssl genrsa -out conf/cert/ca.pem 2048
openssl ecparam -genkey -name secp384r1 -out conf/cert/ca.pem
openssl req -config conf/cert/ca.cnf -newkey rsa:2048 -x509 -days 36500 -key conf/cert/ca.pem -out conf/cert/ca.crt 
```

* ca.cnf
```yaml
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
prompt = no

[req_distinguished_name]   
countryName            = CN
stateOrProvinceName    = Beijing
localityName           = Beijing
postalCode             = 100105
streetAddress          = CNCC, No.7 Tianchen East Road, Chaoyang District
organizationName       = JD Cloud
organizationalUnitName = CPS Dept.
emailAddress           = zhouguiqing3@jd.com
0.commonName           = jdcps 

[v3_req]
keyUsage = keyEncipherment, dataEncipherment
extendedKeyUsage = serverAuth

```

## 2. **Server证书**
* 生成server.key,server.crt (添加了extfile.cnf)
```bash
openssl genrsa -out conf/cert/server.key 2048
openssl ecparam -genkey -name secp384r1 -out conf/cert/server.key
openssl req -config conf/cert/server.cnf -new -key conf/cert/server.key -out conf/cert/server_reqout.txt 
openssl x509 -req -in conf/cert/server_reqout.txt -days 36500 -sha1 -CAcreateserial -CA conf/cert/ca.crt -CAkey conf/cert/ca.pem -out conf/cert/server.crt -extfile conf/cert/extfile.cnf
```

* extfile.cnf
```yaml
subjectAltName = @alt_names
[alt_names]
IP.1 = 127.0.0.1
DNS.1 = localhost
```

* server.cnf
```yaml
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
prompt = no

[req_distinguished_name]
countryName            = CN
stateOrProvinceName    = Beijing
localityName           = Beijing
postalCode             = 100105
streetAddress          = CNCC, No.7 Tianchen East Road, Chaoyang District
organizationName       = JD Cloud
organizationalUnitName = CPS Dept.
emailAddress           = zhouguiqing3@jd.com
0.commonName           = jdcps-proxy.jdcloud.com

[v3_req]
keyUsage = keyEncipherment, dataEncipherment
extendedKeyUsage = serverAuth

```

## 3. **Client证书**
* 生成client.key  client.crt
```bash
openssl genrsa -out conf/cert/client.key 2048
openssl ecparam -genkey -name secp384r1 -out conf/cert/client.key
openssl req -config conf/cert/client.cnf -new -key conf/cert/client.key -out conf/cert/client_reqout.txt 
openssl x509 -req -in conf/cert/client_reqout.txt -days 36500 -sha1 -CAcreateserial -CA conf/cert/ca.crt -CAkey conf/cert/ca.pem -out conf/cert/client.crt
```

* client.cnf
```yaml
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
prompt = no

[req_distinguished_name]
countryName            = CN
stateOrProvinceName    = Beijing
localityName           = Beijing
postalCode             = 100105
streetAddress          = CNCC, No.7 Tianchen East Road, Chaoyang District
organizationName       = JD Cloud
organizationalUnitName = CPS Dept.
emailAddress           = zhouguiqing3@jd.com
0.commonName           = agent 

[ v3_req ]
keyUsage = keyEncipherment, dataEncipherment
extendedKeyUsage = clientAuth

```

### commonName 即CN，ca和client、server的配置中，不能一样