import { describe, it } from "node:test";
import { getFilePathInFilesFolder } from "../../../src/lib/files/utils";
import * as assert from 'node:assert'

describe("utils test suite", () => {
    it("should return joined path", () => {
        const filePath = "test.png"
        const result = getFilePathInFilesFolder(filePath)
        assert.equal(result, "/files/test.png")
    })

    it("should also work if file is defined with ./", () => {
        const filePath = "./test.png"
        const result = getFilePathInFilesFolder(filePath)
        assert.equal(result, "/files/test.png")
    })
})