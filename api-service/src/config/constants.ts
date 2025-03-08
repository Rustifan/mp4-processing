export const statusOptions = [
    'Processing',
    'Failed',
    'Successful',
] as const;

export type FileStatus = typeof statusOptions[number]

export const FILE_FOLDER = "/files"