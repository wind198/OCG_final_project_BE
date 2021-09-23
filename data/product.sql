
INSERT INTO `products` ( `created_at`, `name`, `description`) VALUES (Now(),'King Chart', ' Ngoi suong mong  ');
INSERT INTO `products` ( `created_at`, `name`, `description`) VALUES (Now(),'King Table', 'An suong mom');
INSERT INTO `products` ( `created_at`, `name`, `description`) VALUES (Now(),'King Cup', 'Uong suong mieng');


INSERT INTO `product_variances` (`created_at`,`product_id`, `color`, `size`) VALUES (Now(),'6', 'Red', 'M');
INSERT INTO `product_variances` (`created_at`,`product_id`, `color`, `size`) VALUES (Now(),'6', 'Black', 'M');
INSERT INTO `product_variances` (`created_at`,`product_id`, `color`, `size`) VALUES (Now(),'4', 'Red', 'L');
INSERT INTO `product_variances` (`created_at`,`product_id`, `color`, `size`) VALUES (Now(),'4 ', 'Black', 'L');
INSERT INTO `product_variances` (`created_at`,`product_id`, `color`, `size`) VALUES (Now(),'5', 'Red', 'XL');
INSERT INTO `product_variances` (`created_at`,`product_id`, `color`, `size`) VALUES (Now(),'5', 'Black', 'XL');



INSERT INTO `images` (`created_at`,`product_id`, `image`) VALUES (Now(),'4', 'https://m.media-amazon.com/images/I/41tCIsGV8UL.jpg');
INSERT INTO `images` (`created_at`,`product_id`, `image`) VALUES (Now(),'4', 'https://m.media-amazon.com/images/I/41tCIsGV8UL.jpg');
INSERT INTO `images` (`created_at`,`product_id`, `image`) VALUES (Now(),'5', 'https://www.boconcept.com/on/demandware.static/-/Sites-master-catalog/default/dwffb97acd/images/660000/660620.jpg');
INSERT INTO `images` (`created_at`,`product_id`, `image`) VALUES (Now(),'5', 'data:image/jpeg;base64,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');
INSERT INTO `images` (`created_at`,`product_id`, `image`) VALUES (Now(),'6', 'https://i.pinimg.com/originals/5b/6a/09/5b6a093aa4b36c2548eb559acbc688f8.jpg');
INSERT INTO `images` (`created_at`,`product_id`, `image`) VALUES (Now(),'6', 'data:image/jpeg;base64,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');


