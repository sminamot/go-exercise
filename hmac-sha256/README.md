以下PHPコードのGo版
```PHP
<?php
$data = 'hogehoge'
for ($i = 0; $i < 100; $i++) {
    $data = hash_hmac('sha256', $data, 'test-key');
}
echo $data;
```
