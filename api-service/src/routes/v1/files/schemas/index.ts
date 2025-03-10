import { Type } from "@sinclair/typebox";

export const errorResponseSchema = Type.Object(
    {
        statusCode: Type.Number({ description: "status code" }),
        error: Type.Number({ description: "error type" }),
        message: Type.String({ description: "Error message" }),
    },
    { description: "Error response" },
);
