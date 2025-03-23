# Debris

**弥合设计库格式之间的鸿沟。**

Debris 是一款跨平台命令行工具，旨在通过在不同的设计库管理应用程序之间实现**双向设计库转换**，从而简化您的设计工作流程。告别与不兼容的库格式作斗争，无缝地在各个平台之间迁移您宝贵的设计资产。

目前，Debris 专注于 **Billfish 和 Pixcall** 库之间的转换，并利用 Eagle 格式作为中间格式。未来的开发将扩展对更多 EDA 格式和直接转换方法的支持。

> **🚧  正在进行中 - 早期开发阶段 🚧**
>
> 请注意，Debris 目前正处于积极开发阶段，**尚未准备好用于生产环境。** 虽然 Billfish 和 Pixcall 库之间（通过 Eagle 作为中间格式）的转换功能已实现，但预计仍可能存在潜在问题和局限性。我们感谢您早期的关注，并鼓励您为它的开发做出贡献！

## 功能

**当前支持：**

* **Billfish ↔ Pixcall 转换（通过 Eagle）：** 将 Billfish 库转换为 Pixcall 库，反之亦然，使用 Eagle 格式作为桥梁。 这使您能够在两个流行的 EDA 工具之间使用库。
* **命令行界面 (CLI)：** 简单直观的命令行界面，易于集成到您的脚本和工作流程中。
* **跨平台兼容性：** 借助 Go 编程语言的跨平台兼容性，Debris 可以在主要的操作系统（macOS、Linux、Windows）上运行。

## 用法

Debris 通过命令行使用。

您需要指定输入和输出格式，以及输入和输出路径。

示例：将 Billfish 库转换为 Eagle 库

```
debris --from billfish --to eagle --input ./Billfish --output ./Eagle
```

基本命令结构为：
```
debris --from <输入格式> --to <输出格式> --input <输入路径> --output <输出路径>
```

* `--from billfish` 或 `-f billfish`：指定输入库格式为 Billfish。
* `--to eagle` 或 `-t eagle`：指定所需的输出库格式为 Eagle。
* `--input ./Billfish` 或 `-i ./Billfish`：指示输入 Billfish 库文件或目录的路径。目前，它期望一个 Billfish 库文件夹。
* `--output ./Eagle` 或 `-o ./Eagle`：指定保存转换后的 Eagle 库的输出目录。
