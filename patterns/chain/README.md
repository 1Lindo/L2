# Цепочка вызовов

Цепочка вызовов — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по
цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать
запрос дальше по цепи.

Применимость:

С помощью Цепочки обязанностей можно связать потенциальных обработчиков в одну цепь и при получении запроса поочерёдно
спрашивать каждого из них, не хочет ли он обработать запрос. Это уместно, когда программа должна обрабатывать
разнообразные запросы
несколькими способами, но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них
понадобятся.

Цепочка обязанностей позволяет запускать обработчиков последовательно один за другим в том порядке, в котором они
находятся в цепочке в ситуации, когда важен строгий порядок выполнения.

В любой момент вы можете вмешаться в существующую цепочку и переназначить связи так, чтобы убрать или добавить новое
звено, когда набор объектов, способных обработать запрос, должен задаваться динамически.

Плюсы + :

1) Уменьшает зависимость между клиентом и обработчиками;
2) Реализует принцип единственной обязанности;
3) Реализует принцип открытости/закрытости.

Минусы - :

1) Запрос может остаться никем не обработанным.

Пример из практики:

1) Обработка HTTP запросов. Например, AuthHandler - проверка авторизации и AccessHandler - предоставление доступа
   представляют обработчики запросов, которые могут
   обрабатывать различные аспекты запроса HTTP. Эти обработчики объединены в цепочку, где AuthHandler проверяет
   авторизацию, а затем передает запрос AccessHandler для проверки доступа к ресурсу.
2) Объединение несколько middleware-обработчиков в цепочку. При каждом запросе все middleware в цепочке будут выполнены
   последовательно. Это позволяет легко добавлять, удалять
   или изменять порядок middleware в приложении.