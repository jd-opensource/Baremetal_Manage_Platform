yum install -y make zlib zlib-devel 
cp -f /usr/lib/rpm/openEuler/config.guess /tmp/Python-2.7.18/config.guess 
cp -f /usr/lib/rpm/openEuler/config.sub /tmp/Python-2.7.18/config.sub 
cd /tmp/Python-2.7.18 
./configure --prefix=/usr/local/python-2.7.18 && make -j 8 && make install
ln -s  /usr/local/python-2.7.18/bin/python   /usr/bin/python 
echo /usr/local/python-2.7.18/lib > /etc/ld.so.conf.d/py2.conf 
ldconfig 
python -m ensurepip