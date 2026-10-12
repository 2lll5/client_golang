# Agent Instructions

## Google Drive Build 交付規範（所有 Agent 必須遵守）
- 所有 Codex/Agent 的 Build 交付目的地：Google Drive `AI Project/client_golang/`；名稱必須與 GitHub Repository 完全相同（AI Project folder ID `1Yhfro6Ss5TSCXCVIZIHEu5h7V1gL8b27`）。`AI Project/` 根目錄只放專案資料夾。
- **專案根目錄直接放最新完整且已驗證的 Build**（ZIP／EXE、SHA-256、分片、必要還原工具）。新版本必須完成 Build／測試，驗證 Drive 檔名、父目錄、檔案大小、checksum 與所有分片，才能升格最新版；缺 ZIP 或只有校驗檔則保留前一完整版本。
- **重大功能、架構、相容性或美術里程碑才另建 `vX.Y.Z-說明/` 子資料夾**；一般舊版移到 `歷史封存/`，不刪除、不隨每次 Build 建版本目錄。不同平台各有當前最新版可並存。
- 新開發均提交／推送 GitHub；昂貴 Actions 只用 `workflow_dispatch` 或正式 `release: published`，不隨每次推送 main 觸發。報告資料夾連結、版本、檔案清單與校驗／驗證狀態。
**English (Codex):** Publish newest complete builds at `AI Project/client_golang/`; save important milestones under `vX.Y.Z-description/`, routine old files in `歷史封存/`, never promote incomplete uploads.
