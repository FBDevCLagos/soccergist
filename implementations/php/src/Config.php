<?php


namespace App;

Class Config {

   
  static $facebookCredentials = [

    'verifyToken' => 'politricks',

    'accessToken' => 'EAALGjd6t0LkBADQ1xV9JCIV9CR77VvDEtOjcshyn4zhnCZBgjamp7kcNVesE6zPiKz0GGZC3jDUeupFQmtb59JhrbKBIQl9v9JOdFmlLZCGXbEZB7t5ZCEZC1St1HbqgVQZCBAkL26K73qzrFAdXsxllr619fx9MYEQr7ZCE7FZABhQjU7dCbAsoM'

  ];

   
  static function getVerifyToken() {

    return  Config::$facebookCredentials['verifyToken'];

  }

  static function getAccessToken() {

    return  Config::$facebookCredentials['accessToken'];

  }

  

}
