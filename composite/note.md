- 组合模式的目的是:让客户端不再区分操作的是组合对象还是叶子对象,而是以一
个统一的方式来操作。
  - 实现这个目标的关键之处,是设计一个抽象的组件类,让它可以代表组合对象和叶
  子对象。这样一来,客户端就不用区分到底操作的是组合对象还是叶子对象了,只需要把它们全部当作组件对象进行统一的操作就可以了。
