RSpec.describe 'start_with and end_with matchers' do 
  describe 'caterpillar' do 
    it 'should check with a substring at the beginning or end' do 
      expect(subject).to start_with('cat')
      expect(subject).to end_with('llar')
    end
  end
  
  describe [:a, :b, :c, :d] do 
    it 'should check elements at begining or end array' do 
      expect(subject).to start_with(:a, :b)
      expect(subject).to end_with(:c, :d)
    end
  end
end
