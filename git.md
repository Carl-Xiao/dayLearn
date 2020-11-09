# git操作

> 用了太多ide导致的结果就是不知道怎么使用git,强制要求自己在平时的使用git原生的命令


- git丢弃本地修改的所有文件
`git checkout .`

- git丢弃单个文件的修改
`git checkout -- readme.txt`
> 命令git checkout -- readme.txt意思就是，把readme.txt文件在工作区的修改全部撤销，这里有两种情况：

> 一种是readme.txt自修改后还没有被放到暂存区，现在，撤销修改就回到和版本库一模一样的状态；

> 一种是readme.txt已经添加到暂存区后，又作了修改，现在，撤销修改就回到添加到暂存区后的状态。