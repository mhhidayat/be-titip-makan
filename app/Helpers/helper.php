<?php

if (! function_exists('res')) {
    function res($message, $data, $statusCode = 200)
    {
        return response()->json([
            'message' => $message,
            'data' => $data
        ], $statusCode);
    }
}
