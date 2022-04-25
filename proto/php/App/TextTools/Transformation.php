<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: textTools.proto

namespace App\TextTools;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>textTools.Transformation</code>
 */
class Transformation extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.textTools.Transformer transformer = 1;</code>
     */
    protected $transformer = 0;
    /**
     * Generated from protobuf field <code>string text = 2;</code>
     */
    protected $text = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $transformer
     *     @type string $text
     * }
     */
    public function __construct($data = NULL) {
        \App\GPBMetadata\TextTools::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.textTools.Transformer transformer = 1;</code>
     * @return int
     */
    public function getTransformer()
    {
        return $this->transformer;
    }

    /**
     * Generated from protobuf field <code>.textTools.Transformer transformer = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setTransformer($var)
    {
        GPBUtil::checkEnum($var, \App\TextTools\Transformer::class);
        $this->transformer = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string text = 2;</code>
     * @return string
     */
    public function getText()
    {
        return $this->text;
    }

    /**
     * Generated from protobuf field <code>string text = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setText($var)
    {
        GPBUtil::checkString($var, True);
        $this->text = $var;

        return $this;
    }

}

