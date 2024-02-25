import React, { useState } from "react";

export const Header: React.FC = () => <h1>React component Headerrrrr</h1>;

export const Body: React.FC = () => {
  const [count, setCount] = useState(0);
  return (
    <div className="prose">
      <h2 className="text-red-400">React component Body</h2>
      <p>Count: {count}</p>
      <button onClick={() => setCount(count + 1)}>Increment</button>
    </div>
  );
};

export const Hello = (name: string) => (
  <div>Hello {name} (Client-side React, rendering server-side data)</div>
);
