'use client';

import { CldUploadWidget } from 'next-cloudinary';
import Image from 'next/image';
import { useCallback } from 'react';

declare global {
  let cloudinary;
}

interface ImageUploadProps {
  setImage: (value: string) => void;
}

const ImageUpload: React.FC<ImageUploadProps> = ({ setImage }) => {
  const handleUpload = useCallback(
    result => {
      setImage(result.info.secure_url);
    },
    [setImage],
  );

  return (
    <CldUploadWidget
      onUpload={handleUpload}
      uploadPreset={process.env.NEXT_PUBLIC_CLOUDINARY_UPLOAD_PRESET}
      options={{ maxFiles: 1 }}
    >
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
