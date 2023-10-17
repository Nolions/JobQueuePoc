# JobQueuePoc 

Job Queue 實作

透過生產者/消費者模式來實作queue

程式啟動後會分別啟動一個透過go routine執行的runner(在此稱為`task job runner`)與一個`http api service`。

`task job runner`扮演消費者的腳色，而http api service中有一個`/task` API則是扮演生產的角色，當呼叫`/task` API就會建立新的任務；然後再透過channel會阻塞的特性來指定一次執行多少個task。
