<?php

namespace QueroFreela;

$isApache2Handler = php_sapi_name() !== 'apache2handler';
$isFcgi = php_sapi_name() !== 'fpm-fcgi';

if (!$isApache2Handler && !$isFcgi) {
    die('This script is meant to be run under apache2handler');
}

require_once __DIR__ . '/vendor/autoload.php';

$req = new \QueroFreela\WebRequest();
$res = $req->run();
$res->write();
