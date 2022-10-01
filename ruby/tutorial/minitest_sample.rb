require 'minitest/autorun'

class MinitestSample < Minitest::Test
  def test_sample
    assert_equal 'RUBY', 'ruby'.upcase
    assert_equal 'RUBY', nil.upcase
  end
end
