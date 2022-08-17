# coding: utf-8

"""
    openapi 3.0.3 sample spec

    sample spec for testing openapi functionality, built from json schema tests for draft6  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by: https://openapi-generator.tech
"""

import unittest

import unit_test_api
from unit_test_api.model.minproperties_validation import MinpropertiesValidation
from unit_test_api import configuration


class TestMinpropertiesValidation(unittest.TestCase):
    """MinpropertiesValidation unit test stubs"""
    _configuration = configuration.Configuration()

    def test_ignores_arrays_passes(self):
        # ignores arrays
        MinpropertiesValidation._from_openapi_data(
            [
            ],
            _configuration=self._configuration
        )

    def test_ignores_other_non_objects_passes(self):
        # ignores other non-objects
        MinpropertiesValidation._from_openapi_data(
            12,
            _configuration=self._configuration
        )

    def test_too_short_is_invalid_fails(self):
        # too short is invalid
        with self.assertRaises((unit_test_api.ApiValueError, unit_test_api.ApiTypeError)):
            MinpropertiesValidation._from_openapi_data(
                {
                },
                _configuration=self._configuration
            )

    def test_ignores_strings_passes(self):
        # ignores strings
        MinpropertiesValidation._from_openapi_data(
            "",
            _configuration=self._configuration
        )

    def test_longer_is_valid_passes(self):
        # longer is valid
        MinpropertiesValidation._from_openapi_data(
            {
                "foo":
                    1,
                "bar":
                    2,
            },
            _configuration=self._configuration
        )

    def test_exact_length_is_valid_passes(self):
        # exact length is valid
        MinpropertiesValidation._from_openapi_data(
            {
                "foo":
                    1,
            },
            _configuration=self._configuration
        )


if __name__ == '__main__':
    unittest.main()