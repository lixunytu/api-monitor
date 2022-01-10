package execute

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
)

func KeyExist(arg1, data string) (istrue bool) {
	istrue = true
	strings.Trim(arg1, arg1)
	//首先处理 arg1
	//1.带有 .#. 的需要特殊处理
	split := strings.Split(arg1, ".#.")
	if len(split) == 1 {
		//没有
		result := gjson.Get(data, arg1)
		istrue = result.Exists()
	} else if len(split) == 2 {
		result := gjson.Get(data, split[0])
		if len(result.Array()) == 0 {
			istrue = false
			return
		} else {
			for _, v := range result.Array() {
				result2 := gjson.Get(v.String(), split[1])
				if !result2.Exists() {
					istrue = false
					break
				}
			}
		}

	} else if len(split) == 3 {
		arrystring := ""
		n := len(split)
		for i := 0; i < n-1; i++ {
			arrystring = arrystring + split[i]
			if i != n-2 {
				arrystring = arrystring + ".#."
			}

		}
		result := gjson.Get(data, arrystring)

		if len(result.Array()) == 0 {
			istrue = false
			return
		} else {
			for _, v := range result.Array() {
				//result2 := gjson.Get(v.String(), "#."+split[n-1])
				for _, v2 := range v.Array() {
					get := gjson.Get(v2.String(), split[n-1])
					if !get.Exists() {
						istrue = false
						return
					}
				}
			}
		}
	} else if len(split) == 4 {
		arrystring := ""
		n := len(split)
		for i := 0; i < n-1; i++ {
			arrystring = arrystring + split[i]
			if i != n-2 {
				arrystring = arrystring + ".#."
			}

		}
		result := gjson.Get(data, arrystring)

		if len(result.Array()) == 0 {
			istrue = false
			return
		} else {
			for _, v := range result.Array() {
				//result2 := gjson.Get(v.String(), "#."+split[n-1])
				for _, v2 := range v.Array() {
					for _, v3 := range v2.Array() {
						get := gjson.Get(v3.String(), split[n-1])
						if !get.Exists() {
							istrue = false
							return
						}
					}
				}
			}
		}
	} else if len(split) == 5 {
		arrystring := ""
		n := len(split)
		for i := 0; i < n-1; i++ {
			arrystring = arrystring + split[i]
			if i != n-2 {
				arrystring = arrystring + ".#."
			}

		}
		result := gjson.Get(data, arrystring)

		if len(result.Array()) == 0 {
			istrue = false
			return
		} else {
			for _, v := range result.Array() {
				//result2 := gjson.Get(v.String(), "#."+split[n-1])
				for _, v2 := range v.Array() {
					for _, v3 := range v2.Array() {
						for _, v4 := range v3.Array() {
							get := gjson.Get(v4.String(), split[n-1])
							if !get.Exists() {
								istrue = false
								return
							}
						}
					}
				}
			}
		}
	}
	return
}

func DataCount(arg1, arg2, data string) (istrue bool, err error, errstring string) {
	//errnum 1 是 arg2 不对,0成功,2数值不对
	//counttype 0 是区间,1 是具体数值
	istrue = true
	counttype := 0
	strings.Trim(arg1, arg1)
	strings.Trim(arg2, arg2)
	if arg1[len(arg1)-1] != 35 {
		judge := gjson.Get(data, arg1)
		if !judge.IsArray() {
			return false, nil, "value is not Array"
		}

	} else {
		split := strings.Split(arg1, ".#")
		judge := gjson.Get(data, split[0])
		if !judge.IsArray() {
			return false, nil, "value is not Array"
		}
	}
	//首先判断最后一个字符是不是#
	if arg1[len(arg1)-1] != 35 {
		arg1 = arg1 + ".#"
	}


	result := gjson.Get(data, arg1)


	if result.Type.String() != "Number" {
		return false, nil, "执行结果:" + result.String() + ",期待结果:" + arg2
	}
	//判断 arg2是一个具体的数还是一个区间
	//区间
	if arg2[0] == 91 && arg2[len(arg2)-1] == 93 {
		counttype = 1
	}
	if counttype == 0 {
		//具体数值
		arg22, err := strconv.ParseInt(arg2, 10, 64)
		if err != nil {
			istrue = false
			return istrue, err, "参数格式错误"
		}
		if int(result.Num) == int(arg22) {
			istrue = true
			return istrue, nil, "执行结果:" + strconv.Itoa(int(result.Num)) + ",期待结果:" + arg2
		} else {
			istrue = false
			return istrue, nil, "执行结果:" + strconv.Itoa(int(result.Num)) + ",期待结果:" + arg2
		}
	} else {
		//区间,首先要获取初始变量和结束变量

		//切割[x-x]为 x1-x2
		arg2 = arg2[1 : len(arg2)-1]
		//获取x1,x2 字符串
		float_split := strings.Split(arg2, "-")
		if len(float_split) == 1 {
			//具体数值
			arg22, err := strconv.ParseInt(float_split[0], 10, 64)
			if err != nil {
				istrue = false
				return istrue, err, "参数格式错误"
			}
			if int(result.Num) == int(arg22) {
				istrue = true
				return istrue, nil, "执行结果:" + strconv.Itoa(int(result.Num)) + ",期待结果:" + float_split[0]
			} else {
				istrue = false
				return istrue, nil, "执行结果:" + strconv.Itoa(int(result.Num)) + ",期待结果:" + float_split[0]
			}
		} else {
			arg2_1, err := strconv.ParseInt(float_split[0], 10, 64)
			if err != nil {
				fmt.Println(err)
				istrue = false
				return istrue, err, "参数格式错误"
			}
			arg2_2, err := strconv.ParseInt(float_split[1], 10, 64)
			if err != nil {
				fmt.Println(err)
				istrue = false
				return istrue, err, "参数格式错误"
			}
			if int(result.Num) >= int(arg2_1) && int(result.Num) <= int(arg2_2) {
				istrue = true
				return istrue, nil, "执行结果:" + strconv.Itoa(int(result.Num)) + ",期待结果:" + arg2
			} else {
				istrue = false
				return istrue, nil, "执行结果:" + strconv.Itoa(int(result.Num)) + ",期待结果:" + arg2
			}
		}

	}
}

func Equal(arg1, arg2, data string) (istrue bool, err error, errstring string) {
	istrue = true
	arg1 = strings.TrimSpace(arg1)
	arg2 = strings.TrimSpace(arg2)

	if arg2[0] == 40 && arg2[len(arg2)-1] == 41 {
		//多个情况
		arg2 = arg2[1 : len(arg2)-1]
		arg2_split := strings.Split(arg2, "|")

		result := gjson.Get(data, arg1)
		arg2_split[0] = strings.TrimSpace(arg2_split[0])
		arg2_split[1] = strings.TrimSpace(arg2_split[1])

		if result.IsArray() {
			//如果是数组
			for i, arry := range result.Array() {
				if arry.String() == arg2_split[0] || arry.String() == arg2_split[1] {
					istrue = true
				} else {
					istrue = false
					errstring = "执行结果:下标为:" + strconv.Itoa(i) + "的元素,结果为:" + arry.String() + ",期待结果:" + arg2
					break
				}
			}
			return
		} else {
			//如果不是数组
			//strings.TrimSpace(result.String())
			if result.String() == arg2_split[0] || strings.TrimSpace(result.String()) == arg2_split[1] {
				istrue = true
				errstring = "执行结果:" + strings.TrimSpace(result.String()) + ",期待结果:" + arg2_split[0] + "|" + arg2_split[1]
			} else {
				istrue = false
				errstring = "执行结果:" + strings.TrimSpace(result.String()) + ",期待结果:" + arg2_split[0] + "|" + arg2_split[1]
			}
		}

	} else {
		//单个情况
		result := gjson.Get(data, arg1)
		if result.IsArray() {
			//如果是数组
			istrue = true
			for i, arry := range result.Array() {
				if arry.String() != arg2 {
					errstring = "执行结果:下标为:" + strconv.Itoa(i) + "的元素结果为:" + arry.String() + ",期待结果:" + arg2
					istrue = false
					break
				}
			}
			if istrue {
				errstring = "数组元素为:" + result.String() + ",期待结果为:" + arg2
			}
			return
		} else {
			//如果不是数组
			if strings.TrimSpace(result.String()) == arg2 {
				istrue = true
				errstring = "执行结果:" + result.String() + ",期待结果:" + arg2
			} else {
				istrue = false
				errstring = "执行结果:" + result.String() + ",期待结果:" + arg2
			}
		}
	}

	return
}

func NotNull(arg1, arg2, data string) (istrue bool) {
	if KeyExist(arg1, data) {
		istrue = true
		strings.Trim(arg1, arg1)
		//首先处理 arg1
		//1.带有 .#. 的需要特殊处理
		split := strings.Split(arg1, ".#.")
		if len(split) == 1 {
			//没有
			result := gjson.Get(data, arg1)
			if result.String() == "" || result.String() == "null" || result.String() == "undefined" {
				istrue = false
			} else {
				istrue = true
			}

		} else {
			result := gjson.Get(data, split[0])
			for _, v := range result.Array() {
				result2 := gjson.Get(v.String(), split[1])
				if result2.String() == "" || result2.String() == "null" || result2.String() == "undefined" {
					istrue = false
					break
				}
			}
		}
	} else {
		istrue = false
	}

	return
}

func NotEqual(arg1, arg2, data string) (istrue bool, err error, errstring string) {
	istrue = true
	arg1 = strings.TrimSpace(arg1)
	arg2 = strings.TrimSpace(arg2)

	//单个情况
	result := gjson.Get(data, arg1)
	if result.IsArray() {
		//如果是数组
		istrue = true
		for i, arry := range result.Array() {
			if arry.String() == arg2 {
				errstring = "执行结果:下标为:" + strconv.Itoa(i) + "的元素结果为:" + arry.String() + ",期待结果:" + arg2
				istrue = false
				break
			}
		}
		if istrue {
			errstring = "数组元素为:" + result.String() + ",期待结果为:" + arg2
		}
		return
	} else {
		//如果不是数组
		if strings.TrimSpace(result.String()) != arg2 {
			istrue = true
			errstring = "执行结果:" + result.String() + ",期待结果:" + arg2
		} else {
			istrue = false
			errstring = "执行结果:" + result.String() + ",期待结果:" + arg2
		}
	}
	//}

	return
}

func Include(arg1, arg2, data string) (istrue bool, err error, errstring string) {
	istrue = true
	arg1 = strings.TrimSpace(arg1)
	arg2 = strings.TrimSpace(arg2)

	//单个情况
	result := gjson.Get(data, arg1)
	if result.IsArray() {
		//如果是数组
		istrue = false
		for i, arry := range result.Array() {
			if arry.String() == arg2 {
				errstring = "执行结果:下标为:" + strconv.Itoa(i) + "的元素结果为:" + arry.String() + ",期待结果:" + arg2
				istrue = true
				break
			}
		}
		if istrue {
			errstring = "数组元素为:" + result.String() + ",期待结果为:" + arg2
		} else {
			errstring = "数组中未包含 " + arg2
		}
		return
	} else {
		//如果不是数组
		istrue = false
		errstring = "输入的 key 不是数组"
	}
	//}
	return
}
func GTF(arg1, arg2, data string) (istrue bool, err error, errstring string) {
	istrue = true
	arg1 = strings.TrimSpace(arg1)
	arg2 = strings.TrimSpace(arg2)


	result := gjson.Get(data, arg1)

	//如果两个参数都是数字
	if IsNum(strings.TrimSpace(result.String())) && IsNum(arg2) {

		result_num, _ := strconv.ParseFloat(strings.TrimSpace(result.String()), 10)

		arg2_num, _ := strconv.ParseFloat(strings.TrimSpace(arg2), 10)

		if result_num > arg2_num {
			istrue = true
			errstring = "执行结果:" + result.String() + ",期待结果:" + arg2
		} else {
			istrue = false
			errstring = "执行结果:" + result.String() + ",期待结果:" + arg2
		}

	} else {
		istrue = false
		errstring = "参数或阈值不为数字"
	}

	return
}
func LTF(arg1, arg2, data string) (istrue bool, err error, errstring string) {
	istrue = true
	arg1 = strings.TrimSpace(arg1)
	arg2 = strings.TrimSpace(arg2)


	result := gjson.Get(data, arg1)

	//如果两个参数都是数字
	if IsNum(strings.TrimSpace(result.String())) && IsNum(arg2) {

		result_num, _ := strconv.ParseFloat(strings.TrimSpace(result.String()), 10)

		arg2_num, _ := strconv.ParseFloat(strings.TrimSpace(arg2), 10)

		if result_num < arg2_num {
			istrue = true
			errstring = "执行结果:" + result.String() + ",期待结果:" + arg2
		} else {
			istrue = false
			errstring = "执行结果:" + result.String() + ",期待结果:" + arg2
		}

	} else {
		istrue = false
		errstring = "参数或阈值不为数字"
	}

	return
}


func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
