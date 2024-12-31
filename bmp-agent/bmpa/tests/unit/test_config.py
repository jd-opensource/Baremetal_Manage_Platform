import unittest
from unittest.mock import patch


class TestConfig(unittest.TestCase):
    @patch('bmpa.utils.utils.get_agent_params')
    def test_default_values(self, mock_get_agent_params):
        import bmpa.config as config
        mock_get_agent_params.return_value = {}
        self.assertEqual(config.cmdline_path, '/proc/cmdline')
