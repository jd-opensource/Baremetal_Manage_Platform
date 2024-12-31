.PHONY: clean
export CUR_DIR=${PWD}

help:
	@echo "build - build vmlinz initramfs.gz "

install_loongarch64:
	./build.sh -v -d loongarch64_8_4 -b ${CUR_DIR} -i bmp-agent-base:8.4 -a loongarch64

install_ubuntu_2204:
	./build.sh -v -d ubuntu_2204 -b ${CUR_DIR} -i bmp-agent-base:2204 -a x86_64

install_ubuntu_2204_arm64:
	./build.sh -v -d ubuntu_2204_arm64 -b ${CUR_DIR} -i bmp-agent-base:2204 -a arm64

install_ubuntu_2004_arm64:
	./build.sh -v -d ubuntu_2004_arm64 -b ${CUR_DIR} -i bmp-agent-base:2004 -a arm64

install_7_9:
	./build.sh -v -d centos_7_9 -b ${CUR_DIR} -i bmp-agent-base:7.9 -a x86_64

install: install_7_9

install_7_9_arm64:
	./build.sh -v -d centos_7_9_arm64 -b ${CUR_DIR} -i bmp-agent-base:7.9 -a arm64

install_arm: install_7_9_arm64
