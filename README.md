# Stock-Microservice
This is a microservice cutted out of a [Monolith](https://gitlab.com/matthiasstock/monolith).

[![Build Status](https://travis-ci.org/genofire/hs_master-kss-monolith.svg?branch=master)](https://travis-ci.org/genofire/hs_master-kss-monolith)  [![Coverage Status](https://coveralls.io/repos/github/genofire/hs_master-kss-monolith/badge.svg?branch=master)](https://coveralls.io/github/genofire/hs_master-kss-monolith?branch=master)

## Features of this stock mircoservice
* The main functionality of the microservice is to store the goods with their storage location and a time stamp, when they where stored.
* Functionality of the admin frontend
** Add new goods to the stock
** Manually remove a single goods from the stock, for example when they are rancid
** Remove single goods from the stock, when they are send to a costumer
** Block goods from the stock, when a costumer added them to hie cart
* Functionality of the costumer frontend
** Show the store with a traffic light food labelling
** A stock of more then seven goods corresponds to the colour green (sufficient amount of goods)
** A stock between four and seven goods corresponds to the colour orange (moderate amount of goods)
** A stock between zero and three goods corresponds to the colour red (slim amount of goods)
* Optional Features
** Admin frontend: display of a statistic on how many goods where convexed and manually removed from the stock during the last month
** Traffic light food labelling for each good, which indicates whether the good is too old