package function

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
)

// GenerateID 生成唯一ID
func GenerateID() string {
	return uuid.New().String()
}

// GenerateShortID 生成短ID
func GenerateShortID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")[:16]
}

// MD5Hash 计算MD5哈希
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// GenerateRandomString 生成随机字符串
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// IsValidEmail 验证邮箱格式
func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

// IsValidPhone 验证手机号格式
func IsValidPhone(phone string) bool {
	pattern := `^1[3-9]\d{9}$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}

// IsValidIDCard 验证身份证号格式
func IsValidIDCard(idCard string) bool {
	pattern := `^[1-9]\d{5}(18|19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])\d{3}(\d|X|x)$`
	match, _ := regexp.MatchString(pattern, idCard)
	return match
}

// TruncateString 截断字符串
func TruncateString(str string, maxLen int) string {
	if len(str) <= maxLen {
		return str
	}
	return str[:maxLen] + "..."
}

// RemoveHTMLTags 移除HTML标签
func RemoveHTMLTags(html string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(html, "")
}

// ExtractNumbers 提取字符串中的数字
func ExtractNumbers(str string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str, -1)
	var numbers []int
	for _, match := range matches {
		if num, err := strconv.Atoi(match); err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

// ContainsChinese 检查是否包含中文
func ContainsChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

// GetStringLength 获取字符串长度（中文算2个字符）
func GetStringLength(str string) int {
	length := 0
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			length += 2
		} else {
			length += 1
		}
	}
	return length
}

// FormatScore 格式化分数显示
func FormatScore(score int) string {
	if score <= 0 {
		return "暂无"
	}
	return fmt.Sprintf("%d分", score)
}

// FormatRank 格式化位次显示
func FormatRank(rank int) string {
	if rank <= 0 {
		return "暂无"
	}
	return fmt.Sprintf("位次%d", rank)
}

// CalculateScoreLevel 计算分数等级
func CalculateScoreLevel(score int, scoreType int) string {
	// 根据分数类型和分数判断等级
	// 1-文科，2-理科，3-综合改革
	if scoreType == 1 { // 文科
		if score >= 650 {
			return "A+"
		} else if score >= 600 {
			return "A"
		} else if score >= 550 {
			return "B+"
		} else if score >= 500 {
			return "B"
		} else if score >= 450 {
			return "C+"
		} else if score >= 400 {
			return "C"
		} else {
			return "D"
		}
	} else if scoreType == 2 { // 理科
		if score >= 680 {
			return "A+"
		} else if score >= 630 {
			return "A"
		} else if score >= 580 {
			return "B+"
		} else if score >= 530 {
			return "B"
		} else if score >= 480 {
			return "C+"
		} else if score >= 430 {
			return "C"
		} else {
			return "D"
		}
	} else { // 综合改革或其他
		if score >= 650 {
			return "A+"
		} else if score >= 600 {
			return "A"
		} else if score >= 550 {
			return "B+"
		} else if score >= 500 {
			return "B"
		} else if score >= 450 {
			return "C+"
		} else if score >= 400 {
			return "C"
		} else {
			return "D"
		}
	}
}

// ParseSubjects 解析科目字符串
func ParseSubjects(subjectsStr string) []string {
	if subjectsStr == "" {
		return []string{}
	}
	// 按逗号、空格、顿号等分隔符分割
	re := regexp.MustCompile(`[,，\s]+`)
	subjects := re.Split(subjectsStr, -1)
	var result []string
	for _, subject := range subjects {
		subject = strings.TrimSpace(subject)
		if subject != "" {
			result = append(result, subject)
		}
	}
	return result
}

// FormatSubjects 格式化科目显示
func FormatSubjects(subjects []string) string {
	if len(subjects) == 0 {
		return "未设置"
	}
	return strings.Join(subjects, "、")
}

// CalculateAge 根据出生年份计算年龄
func CalculateAge(birthYear int) int {
	currentYear := time.Now().Year()
	if birthYear <= 0 || birthYear > currentYear {
		return 0
	}
	return currentYear - birthYear
}

// IsValidYear 验证年份
func IsValidYear(year int) bool {
	currentYear := time.Now().Year()
	return year >= 2000 && year <= currentYear+1
}

// GetCurrentYear 获取当前年份
func GetCurrentYear() int {
	return time.Now().Year()
}

// GetAdmissionYears 获取历年数据年份列表
func GetAdmissionYears() []int {
	currentYear := time.Now().Year()
	var years []int
	for i := 0; i < 5; i++ { // 过去5年的数据
		years = append(years, currentYear-i)
	}
	return years
}

// SafeIntConvert 安全转换字符串到int
func SafeIntConvert(str string, defaultValue int) int {
	if str == "" {
		return defaultValue
	}
	if value, err := strconv.Atoi(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeFloatConvert 安全转换字符串到float64
func SafeFloatConvert(str string, defaultValue float64) float64 {
	if str == "" {
		return defaultValue
	}
	if value, err := strconv.ParseFloat(str, 64); err == nil {
		return value
	}
	return defaultValue
}

// SplitAndTrim 分割并修剪字符串
func SplitAndTrim(str, sep string) []string {
	if str == "" {
		return []string{}
	}
	parts := strings.Split(str, sep)
	var result []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// JoinStrings 连接字符串数组
func JoinStrings(items []string, sep string) string {
	return strings.Join(items, sep)
}

// IsEmptyString 检查字符串是否为空
func IsEmptyString(str string) bool {
	return strings.TrimSpace(str) == ""
}

// Capitalize 首字母大写
func Capitalize(str string) string {
	if str == "" {
		return str
	}
	return strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
}

// MaskPhoneNumber 遮罩手机号
func MaskPhoneNumber(phone string) string {
	if len(phone) != 11 {
		return phone
	}
	return phone[:3] + "****" + phone[7:]
}

// MaskIDCard 遮罩身份证号
func MaskIDCard(idCard string) string {
	if len(idCard) < 8 {
		return idCard
	}
	return idCard[:4] + strings.Repeat("*", len(idCard)-8) + idCard[len(idCard)-4:]
}

