# ecommerce-search
电商搜索服务

# TODO
- 服务搭建 gin+es
- es.商品设计
- 评分设计



# es库表结构
spuinfo
商品表的结构
    spu_id                '商品id',
    spu_type              商品类型(关联t_spu_type)',
    goods_name            商品名称',
    price                '商品价格',
    visit_num             访问量',
    goods_tag             '商品标签',
    sale_status           '上架状态： 0下架 1上架 2（定时上架还未上架阶段）待上架',
    appraise_num          '评价数',
    is_deleted            '0正常 1已删除',
    created_at            '创建时间',
    updated_at            '更新时间，有修改自动更新',
    
```
{
  "spu_id": 12345,
  "spu_type": "服装",
  "goods_name": "随机商品名称",
  "price_low": 1299,
  "price_line": 1899,
  "visit_num": 10000,
  "goods_tag": ["热销", "新品"],
  "sale_status": 1,
  "appraise_num": 567,
  "is_deleted": 0,
  "created_at": "2022-10-21 12:34:56",
  "updated_at": "2023-04-17 09:45:23"
}    
```