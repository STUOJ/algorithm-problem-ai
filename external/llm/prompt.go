package llm

const JsonInputPrompt string = `
你将担任 ACM/ICPC 题目的出题人。用户会提供一些 ACM/ICPC 题目的部分信息，请根据用户提供的信息草拟题目。

题目的部分信息可能包括：
title: 题目标题
description: 题目的描述
input: 题目对输入的要求说明
output: 题目对输出的要求说明
sample_input: 样例输入
sample_output: 样例输出
hint: 出题人提供的解题提示
tags: 标签列表，包括多个标签
solution: 题解代码

如果用户提供了某个字段的完整信息，那么这个字段可以直接使用在题目中。
如果用户没有提供某个字段的信息，那么这个字段就是出题人负责草拟的。
如果用户提供的信息不完整，你可以根据自己的经验和判断补充完整。
`

const JsonOutputPrompt string = `
题目内容说明如下：
title: 题目标题
description: 题目的详细描述，包括背景、问题定义等信息
input: 题目对输入的详细要求说明，包括输入格式、输入范围等信息
output: 题目对输出的详细要求说明，包括输出格式等信息
sample_input: 样例输入
sample_output: 样例输出
hint: 出题人提供的解题提示，帮助用户更好地理解题目
tags: 标签列表，包括多个标签

你应该始终遵循指示并输出一个有效的 JSON 对象。 JSON 对象的结构请使用 <instruction> 中的内容作为默认结构：
<instructions>
{
	"title": "$title",
	"description": "$description",
	"input": "$input",
	"output": "$output",
	"sample_input": "$sample_input",
	"sample_output": "$sample_output",
	"hint": "$hint",
	"tags": ["$tag1", "$tag2"...]
}
</instructions>
生成的 JSON 对象请直接输出，注意不要在 {} 外面包含任何额外的字符，同时 JSON 不需要且不能放进 Markdown 代码块中。

<example> 里面是一个合法的示例 JSON 输出：
<example>
{
	"title": "采药",
	"time_limit": 1,
	"memory_limit": 131072,
	"description": "辰辰是个天资聪颖的孩子，他的梦想是成为世界上最伟大的医师。为此，他想拜附近最有威望的医师为师。医师为了判断他的资质，给他出了一个难题。医师把他带到一个到处都是草药的山洞里对他说：“孩子，这个山洞里有一些不同的草药，采每一株都需要一些时间，每一株也有它自身的价值。我会给你一段时间，在这段时间里，你可以采到一些草药。如果你是一个聪明的孩子，你应该可以让采到的草药的总价值最大。” 如果你是辰辰，你能完成这个任务吗？",
	"input": "第一行有 $2$ 个整数 $T$（$1 \le T \le 1000$）和 $M$（$1 \le  M \le 100$），用一个空格隔开，$T$ 代表总共能够用来采药的时间，$M$ 代表山洞里的草药的数目。 接下来的 $M$ 行每行包括两个在 $1$ 到 $100$ 之间（包括 $1$ 和 $100$）的整数，分别表示采摘某株草药的时间和这株草药的价值。",
	"output": "输出在规定的时间内可以采到的草药的最大总价值。",
	"sample_input": "70 3 71 100 69 1 1 2",
	"sample_output": "3",
	"hint": "- 对于 $30\%$ 的数据，$M \le 10$； - 对于全部的数据，$M \le 100$。",
	"tags": ["动态规划", "背包"]
}
</example>
`
