### epoll

epoll底层实现最重要的两个数据结构：epitem和eventpoll。
```
struct epitem {
  struct rb_node  rbn;      
  struct list_head  rdllink; 
  struct epitem  *next;      
  struct epoll_filefd  ffd;  
  int  nwait;                 
  struct list_head  pwqlist;  
  struct eventpoll  *ep;      
  struct list_head  fllink;   
  struct epoll_event  event;  
};

struct eventpoll {
  spin_lock_t       lock; 
  struct mutex      mtx;  
  wait_queue_head_t     wq; 
  wait_queue_head_t   poll_wait; 
  struct list_head    rdllist;   //就绪链表
  struct rb_root      rbr;      //红黑树根节点 
  struct epitem      *ovflist;
};
```

api：
```//用户数据载体
typedef union epoll_data {
   void    *ptr;
   int      fd;
   uint32_t u32;
   uint64_t u64;
} epoll_data_t;
//fd装载入内核的载体
 struct epoll_event {
     uint32_t     events;    /* Epoll events */
     epoll_data_t data;      /* User data variable */
 };
 //api
int epoll_create(int size); 
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);  
int epoll_wait(int epfd, struct epoll_event *events,
                 int maxevents, int timeout);
```
+ poll_create是在内核区创建一个epoll相关的一些列结构，并且将一个句柄fd返回给用户态，后续的操作都是基于此fd的，参数size是告诉内核这个结构的元素的大小，类似于stl的vector动态数组，如果size不合适会涉及复制扩容，不过貌似4.1.2内核之后size已经没有太大用途了；

+ epoll_ctl是将fd添加/删除于epoll_create返回的epfd中，其中epoll_event是用户态和内核态交互的结构，定义了用户态关心的事件类型和触发时数据的载体epoll_data；

+ epoll_wait*是阻塞等待内核返回的可读写事件，epfd还是epoll_create的返回值，events是个结构体数组指针存储epoll_event，也就是将内核返回的待处理epoll_event结构都存储下来，maxevents告诉内核本次返回的最大fd数量，这个和events指向的数组是相关的；

+ epoll_wait是用户态需监控fd的代言人，后续用户程序对fd的操作都是基于此结构的。