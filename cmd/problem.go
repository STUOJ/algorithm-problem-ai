package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"neko-acm/internal/model"
	"neko-acm/internal/service/problem"
	"os"
	"strings"
)

// 生成题目
var ProblemCmd = &cobra.Command{
	Use:   "problem",
	Short: "Generate a problem.",
	Long:  "Generate an ACM-ICPC algorithm problem.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(" -------- 生成题目 -------- ")
		reader := bufio.NewReader(os.Stdin)
		pi, err := readProblemInstruction(reader)
		if err != nil {
			return err
		}

		for {
			// 生成题目
			fmt.Println("正在生成题目...")
			p, err := problem.Generate(pi)
			if err != nil {
				log.Println(err)

				// 重试
				err := clearBuffer(reader)
				if err != nil {
					return err
				}
				fmt.Print("生成失败，是否重试(Y/N)?")
				again, _ := reader.ReadString('\n')
				again = strings.TrimSpace(again)
				again = strings.ToLower(again)
				if again != "y" {
					break
				}
				continue
			}
			fmt.Println("题目生成成功")

			// 保存到文件
			err = saveProblemJson(reader, p)
			if err != nil {
				log.Println(err)
			}

			// 继续生成
			_, err = reader.Discard(reader.Buffered())
			if err != nil {
				return err
			}
			fmt.Print("是否继续生成题目(Y/N)?")
			again, _ := reader.ReadString('\n')
			again = strings.TrimSpace(again)
			again = strings.ToLower(again)
			if again != "y" {
				break
			}
		}

		return nil
	},
}

// 读取题目信息
func readProblemInstruction(reader *bufio.Reader) (model.ProblemInstruction, error) {
	pi := model.ProblemInstruction{}
	err := clearBuffer(reader)
	if err != nil {
		return model.ProblemInstruction{}, err
	}

	fmt.Println("请输入题目信息：")
	fmt.Print("标题：")
	pi.Title, _ = reader.ReadString('\n')
	pi.Title = strings.TrimSpace(pi.Title)
	fmt.Print("描述：")
	pi.Description, _ = reader.ReadString('\n')
	pi.Description = strings.TrimSpace(pi.Description)
	fmt.Print("输入说明：")
	pi.Input, _ = reader.ReadString('\n')
	pi.Input = strings.TrimSpace(pi.Input)
	fmt.Print("输出说明：")
	pi.Output, _ = reader.ReadString('\n')
	pi.Output = strings.TrimSpace(pi.Output)
	fmt.Print("样例输入：")
	pi.SampleInput, _ = reader.ReadString('\n')
	pi.SampleInput = strings.TrimSpace(pi.SampleInput)
	fmt.Print("样例输出：")
	pi.SampleOutput, _ = reader.ReadString('\n')
	pi.SampleOutput = strings.TrimSpace(pi.SampleOutput)
	fmt.Print("提示：")
	pi.Hint, _ = reader.ReadString('\n')
	pi.Hint = strings.TrimSpace(pi.Hint)
	fmt.Print("标签（以空格分隔）：")
	tagsInput, _ := reader.ReadString('\n')
	pi.Tags = strings.Fields(tagsInput)
	fmt.Print("题解代码：")
	pi.Solution, _ = reader.ReadString('\n')
	pi.Solution = strings.TrimSpace(pi.Solution)

	return pi, nil
}

// 保存题目到文件
func saveProblemJson(reader *bufio.Reader, p model.Problem) error {
	err := clearBuffer(reader)
	if err != nil {
		return err
	}
	fmt.Print("是否保存到文件(Y/N)?")
	save, _ := reader.ReadString('\n')
	save = strings.TrimSpace(save)
	save = strings.ToLower(save)

	if save == "y" {
		path, err := problem.SaveJson(p)
		if err != nil {
			fmt.Println("保存失败！")
			return err
		}
		fmt.Println("保存成功，文件路径：" + path)
	}

	return nil
}
