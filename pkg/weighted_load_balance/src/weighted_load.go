package src

import (
	"fmt"
	"time"
)

type WeightNode struct {
	addr            string
	weight          int // 权重
	currentWeight   int // 当前权重
	effectiveWeight int // 有效权重
	times           int
}

type WeightRoundBalance struct {
	curIndex int
	rss      []*WeightNode
}

func (w *WeightRoundBalance) Add(addr string, weight int) {
	node := &WeightNode{
		addr:            addr,
		weight:          weight,
		effectiveWeight: weight,
	}

	w.rss = append(w.rss, node)
}

func (w *WeightRoundBalance) Next() string {
	total := 0
	var best *WeightNode

	for i := 0; i < len(w.rss); i++ {
		r := w.rss[i]

		// 计算有效权重之和
		total += r.effectiveWeight

		// 变更当前权重 = 当前权重+有效权重
		r.currentWeight += r.effectiveWeight

		// 有效权重默认与权重相同，通讯异常 -1，通讯正常 +1，直到恢复与weight相同
		if r.effectiveWeight < r.weight {
			// 测试通讯
			// r.effectiveWeight--
			// else

			r.effectiveWeight++
		}

		// 选择最大当前权重
		if best == nil || best.currentWeight < r.currentWeight {
			w.curIndex = i
			best = r
		}
	}
	if best == nil {
		return ""
	}
	// 变更当前权重 = 当前权重 - 有效权重之和
	best.currentWeight -= total
	best.times++
	return best.addr
}

func (w *WeightRoundBalance) Get() string {
	return w.Next()
}

func (w *WeightRoundBalance) GetCharsData() []map[string]interface{} {
	var charsData = []map[string]interface{}{
		{
			"name":  "time_line",
			"value": time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	for _, rss := range w.rss {
		item := map[string]interface{}{
			"name":  fmt.Sprintf("%s (%v)", rss.addr, rss.weight),
			"value": rss.currentWeight,
		}
		charsData = append(charsData, item)
	}
	return charsData
}
