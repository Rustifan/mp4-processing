import { describe, it } from "node:test";
import { getFilePathInFilesFolder, getFilePathInProcessedFilesFolder } from "../../../src/lib/files/utils";
import * as assert from 'node:assert'

describe("utils test suite", () => {
    it("should return joined path", () => {
        const filePath = "test.png"
        const result = getFilePathInFilesFolder(filePath)
        assert.equal(result, "/files/test.png")
    })

    it("should also work if file is defined with ./", () => {
        const filePath = "./test.mp4"
        const result = getFilePathInFilesFolder(filePath)
        assert.equal(result, "/files/test.mp4")
    })
    it("should return joined path", () => {
        const filePath = "test.png"
        const result = getFilePathInProcessedFilesFolder(filePath)
        assert.equal(result, "/processed_files/test.png")
    })

    it("should also work if file is defined with ./", () => {
        const filePath = "./test"
        const result = getFilePathInProcessedFilesFolder(filePath)
        assert.equal(result, "/processed_files/test")
    })
})