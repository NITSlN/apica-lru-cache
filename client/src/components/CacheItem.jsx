import React from "react";

const CacheItem = ({ cacheData = [] }) => {
  return (
    <>
      {cacheData?.map(({key, value, expiration}) => (
        <tr key={key}>
          <td className="px-6 py-4 whitespace-nowrap">
            <div className="text-sm text-gray-900">{key}</div>
          </td>
          <td className="px-6 py-4 whitespace-nowrap">
            <div className="text-sm text-gray-900">
              {value}
            </div>
          </td>
          <td className="px-6 py-4 whitespace-nowrap">
            <div className="text-sm text-gray-900">
              {expiration}
            </div>
          </td>
        </tr>
      ))}
    </>
  );
};

export default CacheItem;
