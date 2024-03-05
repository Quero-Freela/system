<?php

namespace QueroFreela;

class WebRequest {
    public $method;
    public $url;
    public $headers;
    public $body;
    public $form;
    public $files;
    public $schema;

    public function __construct() {
        $this->method = $_SERVER['REQUEST_METHOD'];
        $this->url = $_SERVER['REQUEST_URI'];
        $this->schema = $protocol = (!empty($_SERVER['HTTPS']) && $_SERVER['HTTPS'] !== 'off' || $_SERVER['SERVER_PORT'] == 443) ? "https://" : "http://";
        $this->headers = getallheaders();
        $this->headers['REMOTE_ADDR'] = $_SERVER['REMOTE_ADDR'];
        $this->body = file_get_contents('php://input');
        $this->form = $_POST;
        $this->files = $_FILES;

        if (count($this->files) === 0) {
            $this->files = null;
        }

        if (count($this->form) === 0) {
            $this->form = null;
        }

        if (count($this->headers) === 0) {
            $this->headers = null;
        }
    }

    public function serialize() {
        return json_encode($this, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
    }

    /**
     * Run the module with the serialized request
     * @description This method will call the gateway with the serialized request to avoid the segmentation fault
     * @return \QueroFreela\WebResponse
     */
    public function run() {
        $body = $this->serialize();

        $pwd = __DIR__;
        $parentPath = realpath($pwd . "/../");
        $libPath = $parentPath . "/bin/querofreela.so";

        $module = \phpgo_load($libPath, "querofreela");
        if (!$module) {
            throw new \Exception("Module [querofreela] not found");
        }

        return \QueroFreela\WebResponse::deserialize($module->run($body));
    }
}
