<?php
// GENERATED CODE -- DO NOT EDIT!

namespace App\TextTools;

/**
 */
class TextToolsClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * Unary rpc
     * @param \App\TextTools\Transformation $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function TransformText(\App\TextTools\Transformation $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/textTools.TextTools/TransformText',
        $argument,
        ['\App\TextTools\TransformationResult', 'decode'],
        $metadata, $options);
    }

    /**
     * BiDi rpc
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function TransformAndSplitText($metadata = [], $options = []) {
        return $this->_bidiRequest('/textTools.TextTools/TransformAndSplitText',
        ['\App\TextTools\TransformationResult','decode'],
        $metadata, $options);
    }

}
