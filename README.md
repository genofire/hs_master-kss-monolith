# Stock-Microservice
This is a microservice cutted out of a [Monolith](https://gitlab.com/matthiasstock/monolith).

[![Build Status](https://travis-ci.org/genofire/hs_master-kss-monolith.svg?branch=master)](https://travis-ci.org/genofire/hs_master-kss-monolith)  [![Coverage Status](https://coveralls.io/repos/github/genofire/hs_master-kss-monolith/badge.svg?branch=master)](https://coveralls.io/github/genofire/hs_master-kss-monolith?branch=master)

## Features of this stock mircoservice
* save goods via product with timestamp, if it was found and there place in a warehouse (in the admin front end)
  * add new goods
  * manual delete of goods, e.g if the good get fouled
  * remove a good from warehouse, if it was send to customer
  * lock goods, if a customer has it in his bill (and release after time x, if it was not send to customer)
* customer front end
  * display availability with something like a traffic lights

### Nice to have / Can-But-Must-Not-Feature
* display of statistics values like how many goods was removed in a time range
* traffic light in admin front end of goods which get fouled in next time or is already fouled
