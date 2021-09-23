Select Count(id) AS  TotalOrder,SUM(Case when orders.fulfilled_at is not null then orders.total_price else null end) AS TotalSales,
COUNT(Case when orders.fulfilled_at is not null then orders.id else null end) AS PaidOrders,
 Count(Case when orders.fulfilled_at is null then orders.id else null end) as UnpaidOrder from orders where created_at between '2021-09-17' and now()
