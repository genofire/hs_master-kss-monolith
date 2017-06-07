# Stock-Microservice
This is a microservice cutted out of a [Monolith](https://gitlab.com/matthiasstock/monolith).

[![Build Status](https://travis-ci.org/genofire/hs_master-kss-monolith.svg?branch=master)](https://travis-ci.org/genofire/hs_master-kss-monolith)  [![CircleCI](https://circleci.com/gh/genofire/hs_master-kss-monolith/tree/master.svg?style=svg)](https://circleci.com/gh/genofire/hs_master-kss-monolith/tree/master) [![Coverage Status](https://coveralls.io/repos/github/genofire/hs_master-kss-monolith/badge.svg?branch=master)](https://coveralls.io/github/genofire/hs_master-kss-monolith?branch=master) [![GoDoc](https://godoc.org/github.com/genofire/hs_master-kss-monolith?status.svg)](https://godoc.org/github.com/genofire/hs_master-kss-monolith)

## Test of autodeployment

* [Stock-Admin](https://stock.pub.warehost.de/)
* [Easy dummy Shop-Cart in browser-cache](https://stock.pub.warehost.de/dummy_cart/)

## Features of this stock mircoservice
* The main functionality of the microservice is to store the goods with their storage location and the date, when they are too old to be sell.
* Functionality of the admin frontend
  * Add new goods to the stock
  * Manually remove a single goods from the stock, for example when they are rancid
  * Remove single goods from the stock, when they are send to a costumer
  * Block goods from the stock, when a costumer adds them to his cart
* Functionality of the costumer frontend
  * Show the store with a traffic light food labelling
* Optional Features
  * Admin frontend: display of a statistic on the amount and average of goods in the stock
  * Admin frontend: display a traffic light food labelling for each good, which indicates whether the good is too old
