# Порядок заказа такси и поездки

## Заказ такси

Пассажир создаёт поездку

```plantuml
class Trip {
    + Route []Place
    + Price int
    + BabySeat bool
}

class Pace {
    + Lat float
    + Lon float
    + Address string
    + Comment string
}
```

Свободные водители, соответствующие условиям поездки, получают уведомление 
о поездке и могут принять условия пассажира или предложить поменять цену.

Пассажир получает предложения и может принять их или ждать
новых предложений. 

Когда одна из сторон принимает предложение, то вторая получает
запрос подтверждения. После взаимного согласия с условиями поездки
всем свободным водителям рассылается уведомление, что заказ принят,
принявший водитель перестаёт быть свободным.

```plantuml
Passenger -> WS: newTrip
WS -> Driver1: push
WS -> Driver2: push
Driver1 -> WS: changeTrip
WS -> Passenger: push
Passenger -> WS: acceptTrip
WS -> Driver1: push
Driver1 -> WS: acceptTrip
WS ->Passenger: push
WS -> Driver2: push
```

## Посадка пассажира, поездка и оплата

Посадка пассажира, поездка и оплата происходят без поддержки со стороны приложения.
Водитель и пассажир сами договариваются о деталях поездки и оплаты
(например об оплате за ожидание). Водитель и пассажир могут отменить поездку
по своему усмотрению.

Начиная поездку водитель может включить Начало поездки. Это изменит интерфейсы
приложений водителя и пассажира и позволит избежать ошибок. 

##  Завершение поездки

Водитель может завершить поездку. В этом случае пассажир получит
предложение оценить поездку. Завершение поездки сделает водителя
доступным для новых заказов, так же водителю предоставляется возможность
оценить пассажира.

## Ошибки и злоупотребления

Для большинства действий должна быть предусмотрена отмена. Например,
случайное нажатие на Завершить поездку можно отменить и продолжить поездку.

Отмены поездки считаются и используются при расчёте рейтинга водителя
и пассажира.

Водителю и пассажиру доступно изменение оценки в течение некоторого времени
после поездки. Можно поставить 5 по просьбе водителя и позже снизить
ему оценку за просьбу поставить 5.