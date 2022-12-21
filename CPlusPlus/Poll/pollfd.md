### pollfd

pollfd
```
typedef struct pollfd {
        int fd;                         /* 需要被检测或选择的文件描述符*/
        short events;                   /* 对文件描述符fd上感兴趣的事件 */
        short revents;                  /* 文件描述符fd上当前实际发生的事件*/
} pollfd_t;
```
一个pollfd结构体表示一个被监视的文件描述符，通过传递fds[]指示 poll() 监视多个文件描述符，其中：
+ 结构体的events域是监视该文件描述符的事件掩码，由用户来设置这个域。
+ 结构体的revents域是文件描述符的操作结果事件掩码，内核在调用返回时设置这个域。

events域中请求的任何事件都可能在revents域中返回。合法的事件如下：
|  常量   | 说明  |
|  ----  | ----  |
| POLLIN  | 普通或优先级带数据可读 |
| POLLRDNORM  | 普通数据可读 |
| POLLRDBAND  | 优先级带数据可读 |
| POLLPRI  | 高优先级数据可读 |
| POLLOUT  | 普通数据可写 |
| POLLWRNORM  | 普通数据可写 |
| POLLWRBAND  | 优先级带数据可写 |
| POLLERR  | 发生错误 |
| POLLHUP  | 发生挂起 |
| POLLNVAL  | 描述字不是一个打开的文件 |