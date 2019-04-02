<?php
/**
 * Created by IntelliJ IDEA.
 * User: TongZhao
 * Date: 2019/1/6
 * Time: 0:26
 */

$data = array(
    'get' => $_GET,
    'post' => $_POST,
    'server_info' => $_SERVER
);
echo json_encode(
    $data
);