# UCenter 领域设计
## 领域模型
* 租户
* 应用
* 用户
* 角色
* 组织
* 系统
* 日志事件

## 领域模型设计
### 系统
* 设置系统的运维租户，运维租户只能有一个
* 取消租户的系统运维权限，同时取消租户下用户的系统超管标记
* 设置系统的超管用户， 用户只能来自于运维租户
* 取消用户的系统超管标记

### 日志事件
* 添加日志事件
* 查询指定标记的日志清单


### 租户
* 在租户创建用户
* 在租户批量创建用户
* 创建租户, 并获取租户标识
* 在租户创建组织结构
* 设置租户用户的组织信息

### 租户组织
* 创建子组织
* 获取子组织
* 调整组织父节点


### 组织的值对象
* 父节点, 只可以是本租户的组织标记
* 叶子节点表


### 用户
* 设置用户的联系信息(姓名,邮件, 联系方式)
* 设置用户的密码
* 验证用户密码正确性
* 用户名不能重复添加
* 联系方式不能重复添加
* 激活用户
* 禁用用户
* 用户是否激活



### 用户的值对象
* 密码， 不能为空，加密存储； 且不能被返回
* 邮件，不能为空，且格式符合xx@xx
* 联系电话，11位，符合电话格式

### 应用
* 生成应用安全信息
* 重置应用安全信息
* 上线应用的可用功能模块
* 获取可用功能模块列表
* 下线没有被分配过的功能模块
* 修改可用模块的显示名称
* 创建应用的租户实例
* 修改应用的租户实例的可用模块
* 取消应用的租户实例授权
* 为应用的租户实例创建角色实例
* 设置角色实例的可用的功能模块，功能模块只能来源于应用开通给本租户的
* 将已存在租户的用户，导入到应用的租户实例中
* 将已存在租户的用户，批量导入到应用的租户实例中
* 为应用的租户实例创建新用户
* 设应用的租户实例用户的角色，角色只能来自于相同实例
* 为应用的租户实例设置超管用户
* 删除应用的租户实例设置超管用户
* 将用户从应用的租户实例中删除


### 应用的可用模块值对象
* key 在应用内不能重复
* name 可用被修改

### 角色
* 设置角色名称及等级



### 租户聚合方法
* 统计租户的用户数量
* 查询租户的用户清单
* 统计租户已开通的应用实例数据量
* 查询租户开通的应用实例清单
* 查询租户是否为运维租户

### 应用的聚合方法
* 统计使用应用的租户数量
* 查询使用应用的租户实例列表
* 查询应用的租户实例的可用功能列表
* 查询应用的租户实例的角色列表
* 查询应用的租户实例的用户列表
* 查询应用的功能是否被引用
* 查询应用的租户实例的角色是否被引用
* 查询应用的租户实例的超管列表


### 用户的聚合方法
* 查询用户是否为租户实例的超管
* 查询用户在租户实例的超管的功能权限列表
* 查询用户有权限的应用列表
* 验证用户是否有在应用中
* 查询用户是否为系统超管
* 查询用户的联系方式
* 查询用户作为超管的应用实例





