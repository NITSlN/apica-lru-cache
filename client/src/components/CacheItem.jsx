import React from "react";

const CacheItem = ({ cacheData = {}}) => {
  return (
    <>
      {Object.keys(cacheData).map((key) => (
        <tr key={key}>
          <td className="px-6 py-4 whitespace-nowrap">
            <div className="text-sm text-gray-900">{key}</div>
          </td>
          <td className="px-6 py-4 whitespace-nowrap">
            <div className="text-sm text-gray-900">
              {cacheData[key]?.value}
            </div>
          </td>
          <td className="px-6 py-4 whitespace-nowrap">
            <div className="text-sm text-gray-900">
              {cacheData[key]?.expiration}
            </div>
          </td>
        </tr>
      ))}
    </>
  );
};

export default CacheItem;
