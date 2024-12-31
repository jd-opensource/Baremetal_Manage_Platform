package util

import (
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
	"strings"

	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
)

const (
	IP_NETWORK_SPLIT       = "/"
	IP_ADDRESS_REGEX       = `^(((\d{1,2})|(1\d{2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{2})|(2[0-4]\d)|(25[0-5]))$`
	IP_WIDTH         int32 = 32
	IP_MAX_VALUE     int64 = 0xffffffff
	IP_MIN_VALUE     int64 = 0
)

func Gateway(cidr string) string {
	ip, _, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(err.Error())
	}
	IncrIP(ip)
	return ip.String()
}

func Mask(mask net.IPMask) string {
	m, _ := hex.DecodeString(mask.String())
	return net.IP([]byte(m)).String()
}

// IP地址自增
func IncrIP(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}

// IP地址自减
func DecrIP(ip net.IP) {
	length := len(ip)
	for i := length - 1; i >= 0; i-- {
		ip[length-1]--
		if ip[length-1] < 0xFF {
			break
		}
		for j := 1; j < length; j++ {
			ip[length-j-1]--
			if ip[length-j-1] < 0xFF {
				return
			}
		}
	}
}

// RandomMacAddress 生成随机mac地址
func RandomMacAddress() string {

	uuid := commonUtil.GetRandString("", 12, false, false, true)
	first_part := uuid[0:4]
	first_part_num, _ := strconv.ParseInt("0x"+first_part, 16, 32)
	first_hex_string := strconv.FormatInt(first_part_num&0xfeff, 16)
	part := []string{fmt.Sprintf("%04x", first_hex_string), uuid[4:8], uuid[8:12]}
	return strings.Join(part, "-")
}

/*
 * 列举一个CIDR段中的所有IP列表
 */
func GetAllHostIPs(cidr string) []string {
	//不过不包含掩码位，就当做32位IP对待
	if !strings.Contains(cidr, "/") {
		return []string{cidr}
	}
	ip := strings.Split(cidr, "/")[0]
	ipSegs := strings.Split(ip, ".")
	maskLen, _ := strconv.Atoi(strings.Split(cidr, "/")[1])
	if maskLen <= 0 {
		return []string{}
	}
	if maskLen == 32 {
		return []string{ip}
	}
	seg1MinIp, seg1MaxIp := getIpSeg1Range(ipSegs[0], maskLen)
	seg2MinIp, seg2MaxIp := getIpSeg2Range(ipSegs[1], maskLen)
	seg3MinIp, seg3MaxIp := getIpSeg3Range(ipSegs[2], maskLen)
	seg4MinIp, seg4MaxIp := getIpSeg4Range(ipSegs[3], maskLen)

	var ipList []string
	for seg1Ip := seg1MinIp; seg1Ip <= seg1MaxIp; seg1Ip++ {
		for seg2Ip := seg2MinIp; seg2Ip <= seg2MaxIp; seg2Ip++ {
			for seg3Ip := seg3MinIp; seg3Ip <= seg3MaxIp; seg3Ip++ {
				for seg4Ip := seg4MinIp; seg4Ip <= seg4MaxIp; seg4Ip++ {
					oneIp := fmt.Sprintf("%d.%d.%d.%d", seg1Ip, seg2Ip, seg3Ip, seg4Ip)
					ipList = append(ipList, oneIp)
				}
			}
		}
	}
	return ipList
}

/*
 * 对于一个seg1.seg2.seg3.seg4格式的IP，本方法获取seg(ment)3部分的最小值和最大值
 */
func getIpSeg1Range(ipSeg string, maskLen int) (int, int) {
	//如果掩码位>16，此时说明掩码只会影响到seg2，seg3，seg4，seg1直接返回
	if maskLen > 8 {
		segIp, _ := strconv.Atoi(ipSeg)
		return segIp, segIp
	}
	ipSegNum, _ := strconv.Atoi(ipSeg)
	return getIpSegRange(uint8(ipSegNum), uint8(8-maskLen))
}

/*
* 对于一个seg1.seg2.seg3.seg4格式的IP，本方法获取seg(ment)3部分的最小值和最大值
 */
func getIpSeg2Range(ipSeg string, maskLen int) (int, int) {
	//如果掩码位>16，此时说明掩码只会影响到seg3，seg4，seg2直接返回
	if maskLen > 16 {
		segIp, _ := strconv.Atoi(ipSeg)
		return segIp, segIp
	}
	ipSegNum, _ := strconv.Atoi(ipSeg)
	return getIpSegRange(uint8(ipSegNum), uint8(16-maskLen))
}

/*
* 对于一个seg1.seg2.seg3.seg4格式的IP，本方法获取seg(ment)3部分的最小值和最大值
 */
func getIpSeg3Range(ipSeg string, maskLen int) (int, int) {
	//如果掩码位>24，此时说明掩码只会影响到seg4，seg3直接返回
	if maskLen > 24 {
		segIp, _ := strconv.Atoi(ipSeg)
		return segIp, segIp
	}
	ipSegNum, _ := strconv.Atoi(ipSeg)
	return getIpSegRange(uint8(ipSegNum), uint8(24-maskLen))
}

/*
 * 对于一个seg1.seg2.seg3.seg4格式的IP，本方法获取seg(ment)4部分的最小值和最大值
 */
func getIpSeg4Range(ipSeg string, maskLen int) (int, int) {
	ipSegNum, _ := strconv.Atoi(ipSeg)
	segMinIp, segMaxIp := getIpSegRange(uint8(ipSegNum), uint8(32-maskLen))
	if segMinIp == 0 {
		segMinIp = segMinIp + 1 //最后一段的IP不能是0，否则就会变为网段
	}
	return segMinIp, segMaxIp
}

func getIpSegRange(userSegIp, offset uint8) (int, int) {
	var ipSegMax uint8 = 255
	netSegIp := ipSegMax << offset
	segMinIp := netSegIp & userSegIp
	segMaxIp := userSegIp&(255<<offset) | ^(255 << offset)
	return int(segMinIp), int(segMaxIp)
}
