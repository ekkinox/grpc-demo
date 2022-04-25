<?php

namespace App\Controller;

use App\TextTools\TextToolsClient;
use App\TextTools\Transformation;
use App\TextTools\Transformer;
use Grpc\ChannelCredentials;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

class TestController extends AbstractController
{
    #[Route('/test/{value}', name: 'app_test')]
    public function index(string $value): Response
    {

        $client = new TextToolsClient(
            "localhost:50051",
            [
                'credentials' => ChannelCredentials::createInsecure()
            ]
        );


        [$reply, $status] = $client->TransformText(
            new Transformation(
                [
                    'transformer' => Transformer::UPPERCASE,
                    'text' => $value
                ]
            )
        )->wait();

        return $this->json([
            'reply' => $reply->getResult(),
            'status' => $status,
        ]);
    }
}
