'use client';


export default function Interviews() {

  return (
    <div className="mx-auto px-6 max-w-7xl flex flex-col justify-center text-center my-12">
      <div className="flex justify-between items-center mb-5">
        <span className="text-3xl">Interviews</span>
        <button>Set Preferences</button>
      </div>
      <p className="text-3xl font-extrabold my-6">Looking for a practice partner for a mock interview?</p>
      <p className="text-2xl font-semibold my-6">
        Customize your preferences using the preferences button, and hit match to get started!
      </p>
      <button>Match</button>
    </div>
  );
}