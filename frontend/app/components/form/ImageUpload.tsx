'use client';

import { CldUploadWidget } from 'next-cloudinary';
import Image from 'next/image';
import { useCallback } from 'react';

declare global {
  let cloudinary;
}

interface ImageUploadProps {
  onChange: (value: string) => void;
}

const ImageUpload: React.FC<ImageUploadProps> = ({ onChange }) => {
  const handleUpload = useCallback(
    result => {
      onChange(result.info.secure_url);
    },
    [onChange],
  );

  return (
    <CldUploadWidget onUpload={handleUpload} uploadPreset="usohal3c" options={{ maxFiles: 1 }}>
      {({ open }) => {
        return (
          <div
            onClick={() => open?.()}
            className="relative cursor-pointer hover:opacity-70 transition border-dashed border-2 p-20 border-neutral-300 flex flex-col  justify-center items-center gap-4 text-neutral-600"
          >
            <Image src="/imageUpload.svg" alt="upload image" width={100} height={100} />
            <div className="font-semibold text-lg">Click to upload</div>
          </div>
        );
      }}
    </CldUploadWidget>
  );
};
export default ImageUpload;
