/**
 * Example
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { HttpFile } from '../http/http';

export class Cat {
    'hunts'?: boolean;
    'age'?: number;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "hunts",
            "baseName": "hunts",
            "type": "boolean  ",
            "format": ""
        },
        {
            "name": "age",
            "baseName": "age",
            "type": "number  ",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return Cat.attributeTypeMap;
    }

    public constructor() {

    }
}

