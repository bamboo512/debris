# Debris

**彌合設計庫格式之間的鴻溝。**

Debris 是一款跨平台命令列工具，旨在通過在不同的設計庫管理應用程式之間實現**雙向設計庫轉換**，從而簡化您的設計工作流程。告別與不相容的庫格式作鬥爭，無縫地在各個平台之間遷移您寶貴的設計資產。

目前，Debris 專注於 **Billfish 和 Pixcall** 庫之間的轉換，並利用 Eagle 格式作為中間格式。未來的開發將擴展對更多 EDA 格式和直接轉換方法的支援。

> **🚧  正在進行中 - 早期開發階段 🚧**
>
> 請注意，Debris 目前正處於積極開發階段，**尚未準備好用於生產環境。** 雖然 Billfish 和 Pixcall 庫之間（通過 Eagle 作為中間格式）的轉換功能已實現，但預計仍可能存在潛在問題和局限性。我們感謝您早期的關注，並鼓勵您為它的開發做出貢獻！

## 功能

**目前支援：**

* **Billfish ↔ Pixcall 轉換（通過 Eagle）：** 將 Billfish 庫轉換為 Pixcall 庫，反之亦然，使用 Eagle 格式作為橋樑。 這使您能夠在兩個流行的 EDA 工具之間使用庫。
* **命令列介面 (CLI)：** 簡單直觀的命令列介面，易於整合到您的腳本和工作流程中。
* **跨平台相容性：** 借助 Go 程式語言的跨平台相容性，Debris 可以在主要的作業系統（macOS、Linux、Windows）上運行。

## 用法

Debris 通過命令列使用。

您需要指定輸入和輸出格式，以及輸入和輸出路徑。

範例：將 Billfish 庫轉換為 Pixcall 庫

```
debris --from billfish --to pixcall --input ./oldFolder.billfish --output ./newFolder
```

基本命令結構為：
```
debris --from <輸入格式> --to <輸出格式> --input <輸入路徑> --output <輸出路徑>
```

* `--from billfish` 或 `-f billfish`：指定輸入庫格式為 Billfish。
* `--to pixcall` 或 `-t pixcall`：指定所需的輸出庫格式為 Pixcall。
* `--input ./oldFolder.billfish` 或 `-i ./oldFolder.billfish`：指示輸入 Billfish 庫檔案或目錄的路徑。目前，它期望一個 Billfish 庫資料夾。
* `--output ./newFolder` 或 `-o ./newFolder`：指定保存轉換後的 Pixcall 庫的輸出目錄。