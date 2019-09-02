<!--
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-01 17:51:14
 * @LastEditTime: 2019-09-01 18:01:52
 * @LastEditors: Please set LastEditors
 -->
# Gin + Gorilla 聊天室

- 項目立於 2019.8.30

### 2019.9.1

> 目前遇到的問題

- [] Client 和 Server 保持 Keep-Alive (這個在 Gorilla 提供的API中的實現模糊不清)
- [] WebSocket 連接失敗的 close 代碼無法提取,導致連接失敗的錯誤無法判斷