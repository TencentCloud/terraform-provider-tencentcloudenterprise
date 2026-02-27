#!/bin/bash
# 文档生成脚本

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

echo "=== TencentCloud Enterprise Provider 文档生成工具 ==="
echo ""
echo "项目根目录: $PROJECT_ROOT"
echo "文档输出目录: $PROJECT_ROOT/docs"
echo ""

cd "$SCRIPT_DIR"

# 检查 go.mod 是否存在
if [ ! -f "go.mod" ]; then
    echo "❌ go.mod 不存在，请先运行 go mod init"
    exit 1
fi

# 编译工具
echo "⚙️  编译文档生成工具..."
go build -o gendoc . 2>&1 | grep -v "^go: downloading" || true

if [ ! -f "gendoc" ]; then
    echo "❌ 编译失败"
    exit 1
fi

echo "✅ 编译成功"
echo ""

# 运行文档生成
echo "📝 开始生成文档..."
./gendoc

echo ""
echo "✅ 文档生成完成！"
echo ""
echo "生成的文档位于："
echo "  - $PROJECT_ROOT/docs/index.md (Provider 首页)"
echo "  - $PROJECT_ROOT/docs/resources/ (资源文档)"
echo "  - $PROJECT_ROOT/docs/data-sources/ (数据源文档)"
echo ""
echo "提示："
echo "  1. 请检查生成的文档是否正确"
echo "  2. 如有错误，请修改对应的 .go 文件注释"
echo "  3. 确保所有 Schema 都有 Description"
echo "  4. 确保 provider.go 的 Resources List 注释是最新的"
