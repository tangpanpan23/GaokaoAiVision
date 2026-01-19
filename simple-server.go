package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 简化的志愿推荐响应结构
type VolunteerSuggestionResponse struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Categories []SuggestionCategory `json:"categories"`
		AnalysisSummary string `json:"analysis_summary"`
		Recommendations []string `json:"recommendations"`
	} `json:"data,omitempty"`
}

type SuggestionCategory struct {
	Category string `json:"category"`
	Colleges []CollegeSuggestion `json:"colleges"`
	Reason   string `json:"reason"`
}

type CollegeSuggestion struct {
	CollegeCode   string  `json:"college_code"`
	CollegeName   string  `json:"college_name"`
	MajorCode     string  `json:"major_code"`
	MajorName     string  `json:"major_name"`
	Batch         string  `json:"batch"`
	MinScore      int     `json:"min_score"`
	MinRank       int     `json:"min_rank"`
	Year          int     `json:"year"`
	MatchingScore float64 `json:"matching_score"`
	Advantages    string  `json:"advantages"`
	Considerations string `json:"considerations"`
}

// CORS中间件
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// 志愿推荐接口
func volunteerSuggestionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("收到志愿推荐请求")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求参数
	var req struct {
		Province     string   `json:"province"`
		ScoreType    int      `json:"score_type"`
		Score        int      `json:"score"`
		Rank         int      `json:"rank"`
		Subjects     string   `json:"subjects"`
		InterestTags []string `json:"interest_tags"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("解析请求失败: %v\n", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("处理请求: 省份=%s, 分数=%d, 位次=%d\n", req.Province, req.Score, req.Rank)

	// 生成推荐结果
	response := VolunteerSuggestionResponse{
		Code: 200,
		Msg:  "success",
	}

	// 分析概况
	response.Data.AnalysisSummary = fmt.Sprintf("根据你的%s分数%d分、位次%d，结合%s选考科目和%s兴趣，推荐如下志愿方案：",
		getScoreTypeName(req.ScoreType), req.Score, req.Rank, req.Subjects, strings.Join(req.InterestTags, "、"))

	// 冲刺志愿
	response.Data.Categories = append(response.Data.Categories, SuggestionCategory{
		Category: "冲",
		Reason:   "分数有一定竞争力，建议冲刺理想学校",
		Colleges: []CollegeSuggestion{
			{
				CollegeCode:    "10001",
				CollegeName:    "清华大学",
				MajorCode:      "080901",
				MajorName:      "计算机科学与技术",
				Batch:          "一本",
				MinScore:       req.Score + 10,
				MinRank:        req.Rank - 500,
				Year:           2024,
				MatchingScore:  0.85,
				Advantages:     "顶尖计算机专业，师资力量雄厚",
				Considerations: "录取分数线较高，需要全力备考",
			},
			{
				CollegeCode:    "10002",
				CollegeName:    "北京大学",
				MajorCode:      "080902",
				MajorName:      "软件工程",
				Batch:          "一本",
				MinScore:       req.Score + 5,
				MinRank:        req.Rank - 300,
				Year:           2024,
				MatchingScore:  0.80,
				Advantages:     "综合性大学，学科交叉明显",
				Considerations: "专业竞争激烈，建议多手准备",
			},
		},
	})

	// 稳妥志愿
	response.Data.Categories = append(response.Data.Categories, SuggestionCategory{
		Category: "稳",
		Reason:   "分数较为稳定，建议选择有把握的学校",
		Colleges: []CollegeSuggestion{
			{
				CollegeCode:    "10003",
				CollegeName:    "上海交通大学",
				MajorCode:      "080903",
				MajorName:      "信息工程",
				Batch:          "一本",
				MinScore:       req.Score - 5,
				MinRank:        req.Rank + 200,
				Year:           2024,
				MatchingScore:  0.90,
				Advantages:     "工科优势明显，就业前景良好",
				Considerations: "需要保持良好发挥",
			},
		},
	})

	// 保底志愿
	response.Data.Categories = append(response.Data.Categories, SuggestionCategory{
		Category: "保",
		Reason:   "确保录取，建议选择安全系数高的学校",
		Colleges: []CollegeSuggestion{
			{
				CollegeCode:    "10004",
				CollegeName:    "华东师范大学",
				MajorCode:      "040101",
				MajorName:      "教育学",
				Batch:          "一本",
				MinScore:       req.Score - 20,
				MinRank:        req.Rank + 1000,
				Year:           2024,
				MatchingScore:  0.75,
				Advantages:     "师范类专业，就业稳定",
				Considerations: "根据个人职业规划选择",
			},
		},
	})

	// 填报建议
	response.Data.Recommendations = []string{
		"建议按照冲、稳、保的原则合理分配志愿",
		"关注各省高考政策变化和院校调档规则",
		"保持良好心态，认真对待每一次模拟考试",
		"及时关注志愿填报时间节点，避免错过填报时间",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	fmt.Println("志愿推荐响应发送完成")
}

// 获取分数类型名称
func getScoreTypeName(scoreType int) string {
	switch scoreType {
	case 1:
		return "文科"
	case 2:
		return "理科"
	case 3:
		return "综合改革"
	default:
		return "未知"
	}
}

// 用户登录接口
func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("收到用户登录请求")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Code string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("解析登录请求失败: %v\n", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("处理登录: code=%s\n", req.Code)

	// 模拟微信登录验证
	response := map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": map[string]interface{}{
			"user_id":      12345,
			"open_id":      "test_open_id_" + req.Code,
			"token":        "test_jwt_token_" + strconv.FormatInt(time.Now().Unix(), 10),
			"need_profile": true,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	fmt.Println("登录响应发送完成")
}

// 健康检查接口
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("收到健康检查请求")

	response := map[string]interface{}{
		"code":   200,
		"msg":    "OK",
		"status": "healthy",
		"time":   time.Now().Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	fmt.Println("健康检查响应发送完成")
}

func main() {
	// 设置路由
	http.HandleFunc("/volunteer/suggestion", corsMiddleware(volunteerSuggestionHandler))
	http.HandleFunc("/user/login", corsMiddleware(userLoginHandler))
	http.HandleFunc("/health", corsMiddleware(healthHandler))

	fmt.Println("🚀 灯塔志愿小程序后端服务启动")
	fmt.Println("📍 服务地址: http://localhost:8080")
	fmt.Println("📋 可用接口:")
	fmt.Println("   POST /user/login - 用户登录")
	fmt.Println("   POST /volunteer/suggestion - 志愿推荐")
	fmt.Println("   GET /health - 健康检查")
	fmt.Println("💡 模拟服务，数据为测试数据")

	// 尝试不同的端口
	ports := []string{":8080", ":3000", ":5000", ":4000"}
	for _, port := range ports {
		fmt.Printf("尝试启动服务在端口%s...\n", port)
		err := http.ListenAndServe(port, nil)
		if err != nil {
			fmt.Printf("端口%s被占用，尝试下一个端口...\n", port)
		} else {
			break
		}
	}
}