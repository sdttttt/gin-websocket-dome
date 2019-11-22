<!--
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-01 17:51:14
 * @LastEditTime: 2019-09-17 15:49:13
 * @LastEditors: Please set LastEditors
 -->
# Gin + Gorilla 聊天室

- 項目立於 2019.8.30

- 由于 Gin 提供了表单验证的一些特性，我在Controller中就完成了表单的验证工作，这可能看起来有些诡异。

- 为了节省资源，这个项目用到了大量的单例模式（连连接池都是单例的...）
- 在 WebSocket 模块中支持定时的任务列队

### 2019.9.1 <聊天室开发end>

> 功能预期

- [x] 多人实时在线

- [x] 广播任务实时执行

- [x] 细化聊天消息粒度

> 目前遇到的問題

- [] I'm really not good at UI, I might need a friend.

- [x] Client 和 Server 保持 Keep-Alive (這個在 Gorilla 提供的API中的實現模糊不清)
- OK 使用官方给出的例子。根据响应时间重置链接超时时间 <= 2019.9.3

- [x] WebSocket 連接失敗的 close 代碼無法提取,導致連接失敗的錯誤無法判斷
- OK 存在异常直接断开链接，犹豫就会败北~！ <= 2019.9.3

    开发暂时结束 2019.9.3

### 2019.9.17

> 准备开始文件托管的功能开发

> 功能预期

- [] 文件下载上传的托管

- [] 文件分享