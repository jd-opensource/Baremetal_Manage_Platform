#### 导航

# 1：isF---字段，角色管理、用户管理、消息中心（admin）、消息中心（非admin）
# 2：hasAdmin---字段，角色管理、用户管理、消息中心（非admin），用来判断是否有权限
# 3：childrenHasAdmin---消息中心（admin）、个人中心有这个字段【若没有admin权限，则消息中心隐藏，显示一级导航的消息中心，无消息设置】
# 4：isChildrenF---消息中心（admin）、个人中心有这个字段
# 5：path、otherPath，用来判断一级导航、二级导航高亮，若为二级导航，详细配置见children-path&otherPath