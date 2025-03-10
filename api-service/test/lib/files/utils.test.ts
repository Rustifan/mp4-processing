import { describe, it } from "node:test";
import {
    getFilePathInFilesFolder,
    getFilePathInProcessedFilesFolder,
    getVerifiedTypeFromSchemaOrThrow,
} from "../../../src/lib/files/utils";
import * as assert from "node:assert";
import { Type } from "@sinclair/typebox";

describe("utils test suite for joining paths", () => {
    it("should return joined path", () => {
        const filePath = "test.png";
        const result = getFilePathInFilesFolder(filePath);
        assert.equal(result, "/files/test.png");
    });

    it("should also work if file is defined with ./", () => {
        const filePath = "./test.mp4";
        const result = getFilePathInFilesFolder(filePath);
        assert.equal(result, "/files/test.mp4");
    });
    it("should return joined path", () => {
        const filePath = "test.png";
        const result = getFilePathInProcessedFilesFolder(filePath);
        assert.equal(result, "/processed-files/test.png");
    });

    it("should also work if file is defined with ./", () => {
        const filePath = "./test";
        const result = getFilePathInProcessedFilesFolder(filePath);
        assert.equal(result, "/processed-files/test");
    });
});

describe("typebox parser test", () => {
    const schema = Type.Object({
        test: Type.String(),
        other: Type.String(),
    });
    it("should throw error", () => {
        assert.throws(() => {
            const obj = {
                test: "test",
                other: 12,
            };
            getVerifiedTypeFromSchemaOrThrow(obj, schema);
        });
    });
});
